git status -s
if [ -z "$(git status -s)" ]; then
    echo "clean"
    echo "下班"
else
    git add .
    git commit -m "update"
    git push origin master
    echo "下班"
fi
