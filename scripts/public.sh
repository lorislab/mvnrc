
function jsonValue() {
    awk -F"[,:}]" '{for(i=1;i<=NF;i++){if($i~/'$1'\042/){print $(i+1)}}}' | tr -d '"'
}

release_notes="../release/release_notes.md"
version=`git describe --tags --always --dirty=-dev`

echo "# Version ${version}" >> $release_notes
echo "" >> $release_notes

for file in ../release/*.zip; do
    echo 'Upload file: '$file
    link=$(curl --request POST --header "PRIVATE-TOKEN: ${GITLAB_TOKEN}" --form "file=@${file}" https://gitlab.com/api/v4/projects/6939307/uploads | jsonValue "markdown")
    link='* '$link
    echo $link >> $release_notes    
done

#Create a release
curl --request POST --header "PRIVATE-TOKEN: ${GITLAB_TOKEN}" -d @${release_notes} https://gitlab.com/api/v4/projects/6939307/${version}/release