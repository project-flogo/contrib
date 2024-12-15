package soapclient

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"time"

	updater_http "github.com/Updater/http"
	"github.com/Updater/soap"
	soap_http "github.com/Updater/soap/http"
	"github.com/mmussett/mxj"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	if s.XMLMode == false {
		// Set attribure prefix for JSON to XML and XML to JSON conversion
		// e.g If the prefix is set to '@', <car electric="true">Tesla</car> will be converted to {"car": {"@electric": "true", "#text": "Tesla"}}
		mxj.SetAttrPrefix(s.XMLAttributePrefix)
	}

	clientPool := updater_http.NewClientPool()

	if strings.HasPrefix(s.SoapServiceEndpoint, "https") {
		tlsConfig := &tls.Config{}
		tlsConfig.InsecureSkipVerify = true
		if s.EnableTLS {
			if s.ServerCertificate != "" {
				sCert, err := decodeCerts(s.ServerCertificate)
				if err != nil {
					return nil, err
				}
				if sCert != nil {
					caCertPool := x509.NewCertPool()
					pemBlock, _ := pem.Decode(sCert)
					if pemBlock == nil {
						return nil, activity.NewError("Unsupported certificate found. It must be a valid PEM certificate.", "", nil)
					}
					serverCert, err1 := x509.ParseCertificate(pemBlock.Bytes)
					if err1 != nil {
						return nil, err1
					}
					caCertPool.AddCert(serverCert)
					tlsConfig.RootCAs = caCertPool
					tlsConfig.InsecureSkipVerify = false
				}
			}
			if s.ClientCertificate != "" && s.ClientKey != "" {
				cert, err := decodeCerts(s.ClientCertificate)
				if err != nil {
					return nil, err
				}
				key, err := decodeCerts(s.ClientKey)
				if err != nil {
					return nil, err
				}
				if cert != nil && key != nil {
					certificate, err := tls.X509KeyPair(cert, key)
					if err != nil {
						return nil, err
					}
					tlsConfig.Certificates = []tls.Certificate{certificate}
					tlsConfig.ClientAuth = 4
				}
			}
		}
		clientPool.SetDefaultTLSConfig(tlsConfig)
	}

	httpAdapter := soap_http.NewClientAdapter(soap_http.ClientPool(clientPool), soap_http.Timeout(time.Duration(s.Timeout)*time.Second))
	act := &Activity{settings: s}
	act.clientAdapter = httpAdapter
	act.xmlPassThroughMode = s.XMLMode

	return act, nil
}

// Activity is an activity that is used to invoke a SOAP service
type Activity struct {
	settings           *Settings
	clientAdapter      soap_http.ClientAdapter
	xmlPassThroughMode bool
}

func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Invokes a SOAP Operation
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return false, err
	}

	uri := a.settings.SoapServiceEndpoint
	if input.HttpQueryParams != nil {
		uri = uri + "?"
		for k, v := range input.HttpQueryParams {
			uri = uri + k + "=" + v + "&"
		}
		uri = strings.TrimRight(uri, "&")
	}

	soapClient, err := soap.NewClient(uri, soap.SetHTTPClient(a.clientAdapter))
	if err != nil {
		return false, err
	}

	ctx.Logger().Infof("SOAP Service Endpoint: [%s]", uri)
	var env soap.Envelope
	env = &soap.Envelope11{HeaderElem: &soap.Header{}, BodyElem: soap.Body11{}}
	if a.settings.SoapVersion == "1.2" {
		env = &soap.Envelope12{HeaderElem: &soap.Header{}, BodyElem: soap.Body12{}}
	}
	if err != nil {
		ctx.Logger().Errorf("Error building SOAP Envelope: %s", err.Error())
		return false, err
	}

	headers := input.SOAPRequestHeaders
	if headers != nil {
		var headerBytes []byte
		if a.xmlPassThroughMode {
			// Just pass headers as received
			headerBytes = []byte(headers.(string))
		} else {
			h_byets, err := json.Marshal(headers)
			if err != nil {
				ctx.Logger().Errorf("Error marshalling JSON headers: %s", err.Error())
				return false, err
			}
			h_map, err := mxj.NewMapJson(h_byets)
			if err != nil {
				ctx.Logger().Errorf("Error converting JSON to XML headers: %s", err.Error())
				return false, err
			}
			headerBytes, _ = h_map.Xml()
		}
		env.Header().Content = headerBytes
	}

	if input.SOAPRequestBody != nil {
		var b_xml_bytes []byte
		if a.xmlPassThroughMode {
			b_xml_bytes = []byte(input.SOAPRequestBody.(string))
		} else {
			b_bytes, err := json.Marshal(input.SOAPRequestBody)
			if err != nil {
				ctx.Logger().Errorf("Error marshalling JSON body: %s", err.Error())
				return false, err
			}
			b_map, err := mxj.NewMapJson(b_bytes)
			if err != nil {
				ctx.Logger().Errorf("Error converting JSON to XML: %s", err.Error())
				return false, err
			}

			b_xml_bytes, err = b_map.Xml()
			if err != nil {
				ctx.Logger().Errorf("Error reading XML bytes: %s", err.Error())
				return false, err
			}
		}

		if a.settings.SoapVersion == "1.1" {
			env.Body().(*soap.Body11).PayloadElem = b_xml_bytes
		} else {
			env.Body().(*soap.Body12).PayloadElem = b_xml_bytes
		}
	}

	var req *soap.Request
	if input.SoapAction != "" {
		ctx.Logger().Infof("SOAP Action: %s", input.SoapAction)
		req = soap.NewRequest(input.SoapAction, env)
	} else {
		req = &soap.Request{Env: env}
	}

	if ctx.Logger().DebugEnabled() {
		ctx.Logger().Debugf("HTTP Request Headers: %v", req.HTTPHeaders)
		env_bytes, _ := xml.Marshal(env)
		ctx.Logger().Debugf("Request SOAP Envelop: %s", string(env_bytes))
	}

	res, err := soapClient.Do(req)
	if err != nil && res == nil {
		ctx.Logger().Errorf("Error invoking SOAP Service: %s", err.Error())
		return false, err
	}

	output := &Output{HttpStatus: res.StatusCode}
	ctx.Logger().Infof("HTTP Response from Server: %d", res.StatusCode)
	if res.Env == nil {
		ctx.Logger().Error("No SOAP Envelope returned in response. Check the SOAP version configured")
		return false, err
	}

	if res.StatusCode == 200 {
		// Success
		if ctx.Logger().DebugEnabled() {
			env_bytes, _ := xml.Marshal(res.Env)
			ctx.Logger().Debugf("Response SOAP Envelop: %s", string(env_bytes))
		}

		if res.Env.Header() != nil {
			if a.xmlPassThroughMode {
				output.SOAPResponseHeaders = string(res.Env.Header().Content)
			} else {
				headers_xml, err := mxj.NewMapXml(res.Env.Header().Content, false)
				if err != nil {
					return false, err
				}
				output.SOAPResponseHeaders = headers_xml.Old()
			}
		}

		if a.settings.SoapVersion == "1.1" {
			soap11ResponseBody := res.Env.Body().(*soap.Body11)
			if soap11ResponseBody.PayloadElem != nil {
				if a.xmlPassThroughMode {
					output.SOAPResponsePayload = string(soap11ResponseBody.PayloadElem)
				} else {
					mv, err := mxj.NewMapXml(soap11ResponseBody.PayloadElem, false)
					if err != nil {
						return false, err
					}
					output.SOAPResponsePayload = mv.Old()
				}
			}
		} else {
			soap12ResponseBody := res.Env.Body().(*soap.Body12)
			if soap12ResponseBody.PayloadElem != nil {
				if a.xmlPassThroughMode {
					output.SOAPResponsePayload = string(soap12ResponseBody.PayloadElem)
				} else {
					mv, err := mxj.NewMapXml(soap12ResponseBody.PayloadElem, false)
					if err != nil {
						return false, err
					}
					output.SOAPResponsePayload = mv.Old()
				}
			}
		}
	} else if res.StatusCode == 500 {
		// Fault
		if ctx.Logger().DebugEnabled() {
			env_bytes, _ := xml.Marshal(res.Env)
			ctx.Logger().Debugf("Response SOAP Envelop: %s", string(env_bytes))
		}

		if res.Env.Header() != nil {
			if a.xmlPassThroughMode {
				output.SOAPResponseHeaders = string(res.Env.Header().Content)
			} else {
				headers_xml, err := mxj.NewMapXml(res.Env.Header().Content, false)
				if err != nil {
					return false, err
				}
				output.SOAPResponseHeaders = headers_xml.Old()
			}
		}

		if a.settings.SoapVersion == "1.1" {
			soap11ResponseBody := res.Env.Body().(*soap.Body11)
			if soap11ResponseBody.FaultElem != nil {
				output.IsFault = true
				fault := soap11ResponseBody.FaultElem
				if a.xmlPassThroughMode {
					xmlFault, _ := xml.Marshal(soap11ResponseBody.FaultElem)
					output.SOAPResponseFault = string(xmlFault)
				} else {
					faultObj := make(map[string]interface{})
					faultObj["faultcode"] = fault.Code
					faultObj["faultactor"] = fault.Actor
					faultObj["faultstring"] = fault.String
					if fault.Detail != nil && len(fault.Detail.Items) > 0 {
						mv, err := mxj.NewMapXml(fault.Detail.Items, false)
						if err != nil {
							return false, err
						}
						faultObj["detail"] = mv.Old()
					}
					output.SOAPResponseFault = faultObj
				}
			}
		} else {
			soap12ResponseBody := res.Env.Body().(*soap.Body12)
			if soap12ResponseBody.FaultElem != nil {
				output.IsFault = true
				fault := soap12ResponseBody.FaultElem
				if a.xmlPassThroughMode {
					xmlFault, _ := xml.Marshal(soap12ResponseBody.FaultElem)
					output.SOAPResponseFault = string(xmlFault)
				} else {
					faultObj := make(map[string]interface{})
					fc_bytes, err := xml.Marshal(fault.Code)
					if err != nil {
						return false, err
					}
					mv, err := mxj.NewMapXml(fc_bytes, false)
					if err != nil {
						return false, err
					}
					faultObj["Code"] = mv.Old()
					faultObj["Reason"] = fault.Reason
					if fault.Node != "" {
						faultObj["Node"] = fault.Node
					}
					if fault.Role != "" {
						faultObj["Role"] = fault.Role
					}
					if fault.Detail != nil && len(fault.Detail.Items) > 0 {
						mv, err := mxj.NewMapXml(fault.Detail.Items, false)
						if err != nil {
							return false, err
						}
						faultObj["Detail"] = mv.Old()
					}
					output.SOAPResponseFault = faultObj
				}
			}
		}
	}

	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}
	return true, nil
}

func decodeCerts(cert string) ([]byte, error) {
	if cert == "" {
		return nil, fmt.Errorf("Certificate not configured")
	}

	if strings.HasPrefix(cert, "file://") {
		fileName := cert[7:]
		_, err := os.Stat(fileName)
		if err != nil {
			return nil, fmt.Errorf("Certificate file not found: %s", fileName)
		}
		return os.ReadFile(fileName)
	}

	if strings.Contains(cert, "/") || strings.Contains(cert, "\\") {
		_, err := os.Stat(cert)
		if err != nil {
			return nil, fmt.Errorf("Certificate file not found: %s", cert)
		}
		return os.ReadFile(cert)
	}

	// case 5: Attempt to decode as base64
	decode, err := base64.StdEncoding.DecodeString(cert)
	if err != nil {
		return nil, fmt.Errorf("Invalid base64 encoded certificate. Check override value configured to the application property.")
	}
	return decode, nil
}
