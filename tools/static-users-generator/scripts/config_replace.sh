#!/bin/bash

set -e #terminate script immediately in case of errors

TEMPDIR=$(mktemp -d)

cp "/config/src/config.yaml" ${TEMPDIR}

TEMP_FILE_PATH="${TEMPDIR}/config.yaml"
DEST_FILE_PATH="/config/dst/config.yaml"
PLACEHOLDER="#__STATIC_PASSWORDS__"

NEWLINE="%"

NUM=0
for secret in $(kubectl get secrets -l dex-user-config=true --all-namespaces -o json | jq -r -c '.items | .[] | .data')
do

  EMAIL=$(echo -n ${secret} | jq -r '.email')
  if [ "${EMAIL}" == "null" ]
  then
    continue
  fi

  EMAIL=$(echo -n ${EMAIL} | base64 -d)
  
  USERNAME=$(echo -n ${secret} | jq -r '.username')
  if [ "${USERNAME}" == "null" ]
  then
    continue
  fi

  USERNAME=$(echo -n ${USERNAME} | base64 -d)

  PASSWORD=$(echo -n ${secret} | jq -r '.password')

  if [ "${PASSWORD}" == "null" ]
  then
    continue
  fi

  PASSWORD=$(echo -n ${PASSWORD} | base64 -d)

  HASH=$(htpasswd -bnBC 12 "" ${PASSWORD} | tr -d ':\n')

  echo "successfully created user: ${USERNAME}"

  NUM=$(( NUM + 1 ))

  # generate userID
  USER_ID=$(cat /dev/urandom | LC_ALL=C tr -dc 'a-z0-9' | fold -w 32 | head -n 1)
  
  # prepare config map to enable static users
  if [ $NUM -eq 1 ]
  then
    sed -i "s;${PLACEHOLDER};staticPasswords:${NEWLINE}${PLACEHOLDER};g" "${TEMP_FILE_PATH}"
  fi

  VALUE="- email: ${EMAIL}${NEWLINE}  hash: ${HASH}${NEWLINE}  username: ${USERNAME}${NEWLINE}  userID: ${USER_ID}${NEWLINE}${PLACEHOLDER}"
  sed -i "s;${PLACEHOLDER};${VALUE};g" "${TEMP_FILE_PATH}"
done

# remove placeholder
sed -i "s;${PLACEHOLDER};"";g" "${TEMP_FILE_PATH}"

# if there were static users created
if [ $NUM -gt 0 ]
then
  # replace newline placeholders with the real newline and save the configuration file in directory which will be mounted to dex container
  cat ${TEMP_FILE_PATH} | tr "${NEWLINE}" "\n" > ${DEST_FILE_PATH}
else
  # copy the configuration file to directory which will be mounted to dex container
  cp ${TEMP_FILE_PATH} ${DEST_FILE_PATH}
fi