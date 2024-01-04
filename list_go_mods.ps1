$goModFiles = Get-ChildItem -Path "." -Filter "go.mod" -Recurse | Resolve-Path -Relative
$goWorkContent = ""
foreach ($filePath in $goModFiles) {
    $dir = $filePath -replace ".{7}$"
    $goMod = $dir -replace '[\\]', '/'
    $goWorkContent = $goWorkContent + "`n" + $goMod
}
Write-Host $goWorkContent
Write-Host 

