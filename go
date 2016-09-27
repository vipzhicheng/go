#!/bin/bash

# Define config directory and config path
GO_HOME=~
AUTHFILE=$GO_HOME/.go.conf

# Parse options
while getopts gh ARGS
do
  case $ARGS in 
    g)
	extra_options="-D7070"
    shift
	;;
    h)
        echo "usage: go [-gh] foo bar"
        exit 0;
    ;;
    *)
        echo "Unknow option: $ARGS"
        echo "usage: go [-gh] foo bar"
        exit 1;
        ;;
  esac
done

# Config search function
# Support AND logic, separated by blank char
GREP()
{

  if [ -z "$*" ]; then
    var="^[^#].*"
    tmp_var=`cat $AUTHFILE | grep -i $var`
    echo "$tmp_var"
    return
  fi
  
  index=1;
  for var in $@
  do
    if [[ "$var" =~ ^- ]];then
      index=$((index+1));
      continue
    fi

    if [ "$var" = "${!index}" ];then
      var="^[^#].*$var.*"
      tmp_var=`grep -i $var $AUTHFILE`
    else
      var="^[^#].*$var.*"
      tmp_var=`echo "$tmp_var" | grep -i $var`
    fi
  done

  echo "$tmp_var"
}

# Get choice
GET_CHAR() 
{
  read choice
  echo $choice;
}

# Support input keywords by arguments or read from command line input
if [ -z $1 ];then
  echo -n "Please input the server IP or Label： "
  read target
else 
  target=$@
fi

# Parse config search result
count=`GREP $target | wc -l`
targetfullname=`GREP $target | awk '{print $1}' | awk -F ':' '{print $1}'`
port=`GREP $target | awk '{print $1}' | awk -F ':' '{if ($2 > "") {print $2} else {print 22}}'`
passwd=`GREP $target | awk '{print $2}' | awk -F ':' '{if ($2 > "") {print $2} else {print "-"}}'`
sshkey=`GREP $target | awk '{print $2}' | awk -F ':' '{if ($3 > "") {print $3} else {print "-"}}'`
user=`GREP $target | awk '{print $2}' | awk -F ':' '{print $1}'`
label=`GREP $target | awk '{print $3}'`

# Process the case of more than one items in search results.
if [ $count -gt 1 ];then
  echo -e 'Found follow servers: (\033[0;31mWhich one do you want to connect?\033[0m)'

  arrtarget=($targetfullname)
  arruser=($user)
  arrpasswd=($passwd)
  arrlabel=($label)
  arrport=($port)
  arrsshkey=($sshkey)

  length=${#arrtarget[@]}

  for ((i=0; i<$length; i++))
  do
    echo -e '[\033[4;34m'$(($i+1))'\033[0m]' "${arruser[$i]}@${arrtarget[$i]}:${arrport[$i]} ${arrlabel[$i]}"
  done

  # Choose one from search results
  echo -n "Please choose by ID: "
  choice=`GET_CHAR`
  echo ""

  echo $choice;

  if [[ "$choice" =~ ^[0-9]+$ ]]; then
    echo '';
  else
    exit 1;
  fi

  targetfullname=${arrtarget[$(($choice-1))]}
  passwd=${arrpasswd[$(($choice-1))]}
  user=${arruser[$(($choice-1))]}
  label=${arrencoding[$(($choice-1))]}
  port=${arrport[$(($choice-1))]}
  sshkey=${arrsshkey[$(($choice-1))]}
fi

# Bad cases
if [ -z $targetfullname ] || [ -z $user ];then
  echo "No matching server~";
  exit 1;
fi
target=$targetfullname

# Any key value should not be empty
if [ -z $port ]; then
  port=22
fi

if [ -z $passwd ]; then
  passwd=-
fi

if [ -z $extra_options ]; then
  extra_options=-
fi

if [ -z $sshkey ]; then
  sshkey=-
fi

echo "Logging into ${user}@${target} ${label}..."

ssh-expect $user $target $passwd $port $extra_options $sshkey
