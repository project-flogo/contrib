#  List all the components in activity, trigger and function
#  Run tests.

val=true
for i in $(ls -d */*/ | awk '{dir=$0 ;pre="/";home=ENVIRON["PWD"] ; ldir= home pre dir; print ldir }')
do 
   cd $i
   err=$(go test); echo $err
    if [[ $err != *"PASS"* ]] ; then
        val=false
    fi    
done 
if [[ !$val ]]; then
    exit 1
fi 