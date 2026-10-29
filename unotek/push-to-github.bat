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

echo Adding files...
git add .

echo Committing files...
git commit -m "feat: initial commit enterprise clean architecture backend with rbac, migrations, seeders, and swagger" -m "Co-Authored-By: Claude Code <noreply@anthropic.com>"

echo Pushing to origin main...
git push -u origin main

echo Done!
pause
