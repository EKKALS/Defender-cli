package main

import (
    "fmt"
    "Defender-cli/api"
    "strconv"
    "log"
    "os"
)
func main(){
    
    var file string
    if os.Args[2] != ""{
        file = os.Args[2]
    }else{
        log.Fatal("ERRORE NESSUN FILE INSERITO!")
    }
    myapi := api.New{
        Token: "d472816094e275309553b3d9c9d4b42d",
    }
    if os.Args[1] == "-f"{ 
        results := myapi.Scan(file)
        fmt.Println("CONTROLLO MALWARE DEL FILE " + file + ":")
        fmt.Println("Lavasoft ha rilevato: " + strconv.Itoa(results.ScanResults.ScanDetails.Lavasoft.ScanResultI) + " " + results.ScanResults.ScanDetails.Lavasoft.ThreatFound)
        fmt.Println("StopZilla ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.STOPzilla.ScanResultI) + " " + results.ScanResults.ScanDetails.STOPzilla.ThreatFound)
        fmt.Println("Zillya ha rilevato: " + strconv.Itoa(results.ScanResults.ScanDetails.Zillya.ScanResultI) + " " + results.ScanResults.ScanDetails.Zillya.ThreatFound)
        fmt.Println("VirusBlokAda ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.VirusBlokAda.ScanResultI) + " " + results.ScanResults.ScanDetails.VirusBlokAda.ThreatFound)
        fmt.Println("TrendMicro ha rilevato: " + strconv.Itoa(results.ScanResults.ScanDetails.TrendMicro.ScanResultI) + " " + results.ScanResults.ScanDetails.TrendMicro.ThreatFound)
        fmt.Println("SUPERAntiSpyware ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.SUPERAntiSpyware.ScanResultI) + " " + results.ScanResults.ScanDetails.SUPERAntiSpyware.ThreatFound)
        fmt.Println("NProtect ha rilevato: " + strconv.Itoa(results.ScanResults.ScanDetails.NProtect.ScanResultI) + " " + results.ScanResults.ScanDetails.NProtect.ThreatFound)
        fmt.Println("Nanoav ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.Nanoav.ScanResultI) + " " + results.ScanResults.ScanDetails.Nanoav.ThreatFound)
        fmt.Println("FSecure ha rilevato: " + strconv.Itoa(results.ScanResults.ScanDetails.FSecure.ScanResultI) + " " + results.ScanResults.ScanDetails.FSecure.ThreatFound)
        fmt.Println("Eset ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.Eset.ScanResultI) + " " + results.ScanResults.ScanDetails.Eset.ThreatFound)
        fmt.Println("BitDefender ha rilevato: " + strconv.Itoa(results.ScanResults.ScanDetails.BitDefender.ScanResultI) + " " + results.ScanResults.ScanDetails.BitDefender.ThreatFound)
        fmt.Println("Baidu ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.Baidu.ScanResultI) + " " + results.ScanResults.ScanDetails.Baidu.ThreatFound)
        fmt.Println("Ahnlab ha rilevato: " + strconv.Itoa(results.ScanResults.ScanDetails.Ahnlab.ScanResultI) + " " + results.ScanResults.ScanDetails.Ahnlab.ThreatFound)
        fmt.Println("AegisLab ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.AegisLab.ScanResultI) + " " + results.ScanResults.ScanDetails.AegisLab.ThreatFound)
        fmt.Println("Zoner ha rilevato: " + strconv.Itoa(results.ScanResults.ScanDetails.Zoner.ScanResultI) + " " + results.ScanResults.ScanDetails.Zoner.ThreatFound)
        fmt.Println("Sophos ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.Sophos.ScanResultI) + " " + results.ScanResults.ScanDetails.Sophos.ThreatFound)
        fmt.Println("Preventon ha rilevato: " + strconv.Itoa(results.ScanResults.ScanDetails.Preventon.ScanResultI) + " " + results.ScanResults.ScanDetails.Preventon.ThreatFound)
        fmt.Println("McAfee ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.McAfee.ScanResultI) + " " + results.ScanResults.ScanDetails.McAfee.ThreatFound)
        fmt.Println("K7 ha rilevato: " + strconv.Itoa(results.ScanResults.ScanDetails.K7.ScanResultI) + " " + results.ScanResults.ScanDetails.K7.ThreatFound)
        fmt.Println("Jiangmin ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.Jiangmin.ScanResultI) + " " + results.ScanResults.ScanDetails.Jiangmin.ThreatFound)
        fmt.Println("Hauri ha rilevato: " + strconv.Itoa(results.ScanResults.ScanDetails.Hauri.ScanResultI) + " " + results.ScanResults.ScanDetails.Hauri.ThreatFound)
        fmt.Println("FProt ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.FProt.ScanResultI) + " " + results.ScanResults.ScanDetails.FProt.ThreatFound)
        fmt.Println("Fortinet ha rilevato: " + strconv.Itoa(results.ScanResults.ScanDetails.Fortinet.ScanResultI) + " " + results.ScanResults.ScanDetails.Fortinet.ThreatFound)
        fmt.Println("Filseclab ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.Filseclab.ScanResultI) + " " + results.ScanResults.ScanDetails.Filseclab.ThreatFound)
        fmt.Println("Emsisoft ha rilevato: " + strconv.Itoa(results.ScanResults.ScanDetails.Emsisoft.ScanResultI) + " " + results.ScanResults.ScanDetails.Emsisoft.ThreatFound)
        fmt.Println("ClamAV ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.ClamAV.ScanResultI) + " " + results.ScanResults.ScanDetails.ClamAV.ThreatFound)
        fmt.Println("Avira ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.Avira.ScanResultI) + " " + results.ScanResults.ScanDetails.Avira.ThreatFound)
        fmt.Println("Avg ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.Avg.ScanResultI) + " " + results.ScanResults.ScanDetails.Avg.ThreatFound)
        fmt.Println("Agnitum ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.Agnitum.ScanResultI) + " " + results.ScanResults.ScanDetails.Agnitum.ThreatFound)
        fmt.Println("Ikarus ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.Ikarus.ScanResultI) + " " + results.ScanResults.ScanDetails.Ikarus.ThreatFound)
        fmt.Println("Cyren ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.Cyren.ScanResultI) + " " + results.ScanResults.ScanDetails.Cyren.ThreatFound)
        fmt.Println("MicrosoftSecurityEssentials ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.MicrosoftSecurityEssentials.ScanResultI) + " " + results.ScanResults.ScanDetails.MicrosoftSecurityEssentials.ThreatFound)
        fmt.Println("QuickHeal ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.QuickHeal.ScanResultI) + " " + results.ScanResults.ScanDetails.QuickHeal.ThreatFound)
        fmt.Println("TotalDefense ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.TotalDefense.ScanResultI) + " " + results.ScanResults.ScanDetails.TotalDefense.ThreatFound)
        fmt.Println("TrendMicroHouseCall ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.TrendMicroHouseCall.ScanResultI) + " " + results.ScanResults.ScanDetails.TrendMicroHouseCall.ThreatFound)
        fmt.Println("XvirusPersonalGuard ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.XvirusPersonalGuard.ScanResultI) + " " + results.ScanResults.ScanDetails.XvirusPersonalGuard.ThreatFound)
        fmt.Println("DrWebGateway ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.DrWebGateway.ScanResultI) + " " + results.ScanResults.ScanDetails.DrWebGateway.ThreatFound)
        fmt.Println("VirITEXplorer ha rilevato: " + strconv. Itoa(results.ScanResults.ScanDetails.VirITEXplorer.ScanResultI) + " " + results.ScanResults.ScanDetails.VirITEXplorer.ThreatFound)
    }
}
