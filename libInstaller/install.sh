#!/bin/bash

#-----------------------------------------------------------------

echo -e "Achitecture :"
echo -e "1 : x64"
echo -e "2 : x86"
echo -n "Choose : "
read A

ARCH="";

if [ $A -eq 1 ]; then
	ARCH="x64"
else
	if [ $A -eq 2 ]; then
		ARCH="x86"
	fi
fi

if [ $ARCH != "" ]; then

cp $ARCH/libcrypto.so.1.0.0 /usr/lib
ln -s /usr/lib/libcrypto.so.1.0.0 /usr/lib/libcrypto.so

cp $ARCH/libssl.so.1.0.0 /usr/lib
ln -s /usr/lib/libssl.so.1.0.0 /usr/lib/libssl.so

cp $ARCH/libicudata.so.53.1 /usr/lib
ln -s /usr/lib/libicudata.so.53.1 /usr/lib/libicudata.so.53
ln -s /usr/lib/libicudata.so.53.1 /usr/lib/libicudata.so.5
ln -s /usr/lib/libicudata.so.53.1 /usr/lib/libicudata.so

cp $ARCH/libicui18n.so.53.1 /usr/lib
ln -s /usr/lib/libicui18n.so.53.1 /usr/lib/libicui18n.so.53
ln -s /usr/lib/libicui18n.so.53.1 /usr/lib/libicui18n.so.5
ln -s /usr/lib/libicui18n.so.53.1 /usr/lib/libicui18n.so

cp $ARCH/libicuuc.so.53.1 /usr/lib
ln -s /usr/lib/libicuuc.so.53.1 /usr/lib/libicuuc.so.53
ln -s /usr/lib/libicuuc.so.53.1 /usr/lib/libicuuc.so.5
ln -s /usr/lib/libicuuc.so.53.1 /usr/lib/libicuuc.so

cp $ARCH/libQt5Core.so.5.4.1 /usr/lib
ln -s /usr/lib/libQt5Core.so.5.4.1 /usr/lib/libQt5Core.so.5.4
ln -s /usr/lib/libQt5Core.so.5.4.1 /usr/lib/libQt5Core.so.5
ln -s /usr/lib/libQt5Core.so.5.4.1 /usr/lib/libQt5Core.so

cp $ARCH/libQt5Network.so.5.4.1 /usr/lib
ln -s /usr/lib/libQt5Network.so.5.4.1 /usr/lib/libQt5Network.so.5.4
ln -s /usr/lib/libQt5Network.so.5.4.1 /usr/lib/libQt5Network.so.5
ln -s /usr/lib/libQt5Network.so.5.4.1 /usr/lib/libQt5Network.so

cp $ARCH/libQt5Script.so.5.4.1 /usr/lib
ln -s /usr/lib/libQt5Script.so.5.4.1 /usr/lib/libQt5Script.so.5.4
ln -s /usr/lib/libQt5Script.so.5.4.1 /usr/lib/libQt5Script.so.5
ln -s /usr/lib/libQt5Script.so.5.4.1 /usr/lib/libQt5Script.so

cp $ARCH/libQt5Sql.so.5.4.1 /usr/lib
ln -s /usr/lib/libQt5Sql.so.5.4.1 /usr/lib/libQt5Sql.so.5.4
ln -s /usr/lib/libQt5Sql.so.5.4.1 /usr/lib/libQt5Sql.so.5
ln -s /usr/lib/libQt5Sql.so.5.4.1 /usr/lib/libQt5Sql.so

cp  libsqlite3/usr/lib/x86_64-linux-gnu/libsqlite3.so.0.8.6 /usr/lib/x86_64-linux-gnu/
cp  libsqlite3/usr/lib/x86_64-linux-gnu/libsqlite3.so.0 /usr/lib/x86_64-linux-gnu/
ln -s /usr/lib/x86_64-linux-gnu/libsqlite3.so.0 /usr/lib/x86_64-linux-gnu/libsqlite3.so


fi
