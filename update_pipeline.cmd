@echo off
cd /d C:\Users\avasi\AI_WORKER\fastapi_app\site_bundle

sec_fred_update.exe
metrics_clean.exe
build_stock_universe.exe
build_intrinsic_batch.exe

git add data
git commit -m "Automated data update" 
git push
