@echo off
echo ==============================================
echo Pushing Unotek Backend to GitHub...
echo ==============================================

cd /d "%~dp0"

if not exist .git (
    echo Initializing git...
    git init
    git branch -M main
    git remote add origin https://github.com/JackyRahman/unotek.git
)

echo Staging all files...
git add -A

echo Committing files...
git commit -m "feat: restructure backend to repo root, include svg designs, and full clean architecture" -m "Co-Authored-By: Claude Code <noreply@anthropic.com>"

echo Pushing to origin main...
git push -u origin main

echo Done!
pause
