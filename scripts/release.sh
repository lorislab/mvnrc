#!/usr/bin/env bash

output_file=mvnrc

cd ..
rm -rf release
mkdir release

platforms=("linux:amd64" "linux:arm64" "darwin:amd64" "windows:amd64" "windows:386")

for platform in "${platforms[@]}"
do
    platform_split=(${platform/:/ })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_dir='release/builds/'$GOOS'/'$GOARCH'/'
    output_name=$output_dir$output_file
    suffix=''
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
        suffix='.exe'
    fi  
    
    echo 'Build...'$output_name
    env GOOS=$GOOS GOARCH=$GOARCH TAGER_SUFFIX=$suffix TARGET_PREFIX=$output_dir make build
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
    zip_file='release/'$output_file'-'$GOOS'-'$GOARCH'.zip'
    echo 'Package...'$zip_file
    zip -r -j -q $zip_file $output_dir/*
done

cd release/builds
zip_file='../'$output_file'-all.zip'
echo 'Package...'$zip_file
zip -r -q $zip_file ./*
cd ../..