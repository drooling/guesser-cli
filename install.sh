#!/bin/bash

if [ "$EUID" -ne 0 ]
  then echo "This script must be run as root"
  exit
fi

echo "Removing any previous installations..."
rm -f /bin/partialguesser
rm -rf /usr/share/partialguesser

echo "Moving binary to /bin..."
mv ./partialguesser /bin/partialguesser

echo "Creating /usr/share/partialguesser..."
mkdir /usr/share/partialguesser

echo "Fetching domain list..."
wget -q -P /usr/share/partialguesser https://raw.githubusercontent.com/drooling/guesser-cli/data/domains.txt

echo "Installation finished."
echo ""
partialguesser --help