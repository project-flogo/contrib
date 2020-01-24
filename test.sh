#  List all the components in activity, trigger and function
#  Run tests.
setup_kafka () {
    
    cd $1/activity/kafka
    docker-compose up -d
    cd $1
}
destroy_kafka () {
   
    cd $1/activity/kafka
    docker-compose stop
    cd $1
}
pwd=$(pwd)
setup_kafka $pwd

val=true
for i in $(ls -d */*/ | awk '{dir=$0 ;pre="/";home=ENVIRON["PWD"] ; ldir= home pre dir; print ldir }')
do 
   cd $i
   err=$(go test); echo $err
    if [[ $err == *"FAIL"* ]] ; then
        val=false
    fi    
done 

destroy_kafka $pwd

if [[ !$val ]]; then
    exit 1
fi 

