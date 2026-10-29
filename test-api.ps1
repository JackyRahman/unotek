Write-Host "===================================================" -ForegroundColor Cyan
Write-Host "[TEST] Testing Unotek Backend API Endpoints" -ForegroundColor Cyan
Write-Host "===================================================" -ForegroundColor Cyan
Write-Host ""

Write-Host "1. Testing Public Endpoint: GET /api/v1/company/profile" -ForegroundColor Yellow
try {
    $res = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/company/profile" -Method Get
    $res | ConvertTo-Json -Depth 4
} catch {
    Write-Host "Error: $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

Write-Host "2. Testing Super Admin Login: POST /api/v1/auth/login" -ForegroundColor Yellow
$loginBody = @{
    identifier = "admin@company.com"
    password   = "admin12345"
} | ConvertTo-Json

$token = $null
try {
    $loginRes = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/auth/login" -Method Post -Body $loginBody -ContentType "application/json"
    $loginRes | ConvertTo-Json -Depth 4
    $token = $loginRes.data.access_token
} catch {
    Write-Host "Login Error: $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

if ($token) {
    Write-Host "3. Testing Protected Admin Endpoint: GET /api/v1/auth/profile" -ForegroundColor Yellow
    try {
        $headers = @{ Authorization = "Bearer $token" }
        $profileRes = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/auth/profile" -Method Get -Headers $headers
        $profileRes | ConvertTo-Json -Depth 4
    } catch {
        Write-Host "Profile Error: $($_.Exception.Message)" -ForegroundColor Red
    }
    Write-Host ""
}

Write-Host "4. Testing Public Solutions (WMS): GET /api/v1/solutions" -ForegroundColor Yellow
try {
    $solRes = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/solutions" -Method Get
    $solRes | ConvertTo-Json -Depth 4
} catch {
    Write-Host "Solution Error: $($_.Exception.Message)" -ForegroundColor Red
}
Write-Host ""

Write-Host "5. Swagger UI: http://localhost:8080/swagger/index.html" -ForegroundColor Green
Write-Host "===================================================" -ForegroundColor Cyan
