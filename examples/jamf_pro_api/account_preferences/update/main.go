package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/account_preferences"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	authConfig, err := client.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	newSettings := account_preferences.ResourceAccountPreferencesV2{
		Language:                             "en",
		DateFormat:                           "MM/dd/yyyy",
		Timezone:                             "America/Chicago",
		DisableRelativeDates:                 false,
		DisablePageLeaveCheck:                true,
		DisableShortcutsTooltips:             false,
		DisableTablePagination:               true,
		ConfigProfilesSortingMethod:          "ALPHABETICALLY",
		ResultsPerPage:                       20,
		UserInterfaceDisplayTheme:            "DARK",
		ComputerSearchMethod:                 "CONTAINS",
		ComputerApplicationSearchMethod:      "CONTAINS",
		ComputerApplicationUsageSearchMethod: "CONTAINS",
		ComputerFontSearchMethod:             "CONTAINS",
		ComputerPluginSearchMethod:           "CONTAINS",
		ComputerLocalUserAccountSearchMethod: "CONTAINS",
		ComputerSoftwareUpdateSearchMethod:   "CONTAINS",
		ComputerPackageReceiptSearchMethod:   "CONTAINS",
		ComputerPrinterSearchMethod:          "CONTAINS",
		ComputerPeripheralSearchMethod:       "CONTAINS",
		ComputerServiceSearchMethod:          "CONTAINS",
		MobileDeviceSearchMethod:             "CONTAINS",
		MobileDeviceAppSearchMethod:          "CONTAINS",
		UserSearchMethod:                     "CONTAINS",
		UserAllContentSearchMethod:           "CONTAINS",
		UserMobileDeviceAppSearchMethod:      "CONTAINS",
		UserMacAppStoreAppSearchMethod:      "CONTAINS",
		UserEbookSearchMethod:                "CONTAINS",
	}

	updated, _, err := jamfClient.AccountPreferences.UpdateAccountPreferencesV2(context.Background(), &newSettings)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	out, _ := json.MarshalIndent(updated, "", "    ")
	fmt.Println("Updated account preferences:\n" + string(out))
}
