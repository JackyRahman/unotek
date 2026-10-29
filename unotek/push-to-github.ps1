Set-Location $PSScriptRoot

if (-not (Test-Path ".git")) {
    Write-Host "Initializing git repository..." -ForegroundColor Cyan
    git init
    git branch -M main
    git remote add origin https://github.com/JackyRahman/unotek.git
}

Write-Host "Staging all project files..." -ForegroundColor Cyan
git add .

Write-Host "Creating initial commit..." -ForegroundColor Cyan
git commit -m @"
feat: initial commit enterprise clean architecture backend with rbac, migrations, seeders, and swagger

Co-Authored-By: Claude Code <noreply@anthropic.com>
"@

Write-Host "Pushing to remote origin main..." -ForegroundColor Cyan
git push -u origin main

Write-Host "Successfully pushed to https://github.com/JackyRahman/unotek.git" -ForegroundColor Green
