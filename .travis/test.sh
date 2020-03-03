#  List all the components in activity, trigger and function
#  Run tests.
setup_kafka () {
    
    docker-compose up -d
    
   
    
}
setup_pulsar() {
    wget https://archive.apache.org/dist/pulsar/pulsar-2.5.0/apache-pulsar-2.5.0-bin.tar.gz
    tar xvfz apache-pulsar-2.5.0-bin.tar.gz
}
destroy_kafka () {
   cd $1
    docker-compose stop
    
}
pwd=$(pwd)
setup_kafka 
setup_pulsar
val=true

for i in $(ls -d ../*/*/ | awk '{dir=$0 ;pre="/";home=ENVIRON["PWD"] ; ldir= home pre dir; print ldir }')
do 
   cd $i
   err=$(go test); echo $err
    if [[ $err == *"FAIL"* ]] ; then
        val=false
    fi    
done 

destroy_kafka $pwd

case $val in
    (true) exit 0;;
    (false) exit 1;;
esac


