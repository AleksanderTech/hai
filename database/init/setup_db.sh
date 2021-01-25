hai_user=hai
hai_password=hai
hai_database=hai

if [ -z "$1" ]; then
  user='postgres'
else
  user=$1
fi

if [ -z "$2" ]; then
  password='postgres'
else
  password=$2
fi

if [ -z "$3" ]; then
  host='localhost'
else
  host=$3
fi

if [ -z "$4" ]; then
  port=5432
else
  port=$4
fi

psql postgresql://${user}:${password}@${host}:${port} -f init.sql
psql postgresql://${hai_user}:${hai_password}@${host}:${port} -f ../schema.sql
