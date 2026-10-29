Set-Location $PSScriptRoot

if (-not (Test-Path ".git")) {
    Write-Host "Initializing git repository..." -ForegroundColor Cyan
    git init
    git branch -M main
    git remote add origin https://github.com/JackyRahman/unotek.git
}

Write-Host "Staging all project files..." -ForegroundColor Cyan
git add -A

Write-Host "Creating commit..." -ForegroundColor Cyan
git commit -m @"
feat: restructure backend to repo root, include svg designs, and full clean architecture

Co-Authored-By: Claude Code <noreply@anthropic.com>
"@

Write-Host "Pushing to remote origin main..." -ForegroundColor Cyan
git push -u origin main

Write-Host "Successfully pushed to https://github.com/JackyRahman/unotek.git" -ForegroundColor Green
