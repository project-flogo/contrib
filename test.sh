#  List all the components in activity, trigger and function
#  Append './' and run tests.
for i in $(ls -d */*/ | awk '{dir=$0 ;pre="./";ldir= pre dir; print ldir }')
do 
    go test $i
done 