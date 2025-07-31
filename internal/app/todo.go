package app

// func structToMap(config *Cfg) map[string]interface{} {
// 	// Marshal the config struct to JSON
// 	jsonData, err := json.Marshal(config)
// 	if err != nil {
// 		return nil
// 	}
// 	// Decode the JSON data into a map
// 	var configMap map[string]interface{}
// 	err = json.Unmarshal(jsonData, &configMap)
// 	if err != nil {
// 		return nil
// 	}

// 	// Convert MaxUploadSize to a normal string representation
// 	configMap["MaxUploadSize"] = strconv.FormatInt(config.MaxUploadSize, 10)

// 	return configMap
// }

// func saveConfig(configPath string, config *Cfg) error {
// 	data, err := yaml.Marshal(config)
// 	if err != nil {
// 		return fmt.Errorf("failed to marshal config: %w", err)
// 	}

// 	err = os.WriteFile(configPath, data, 0644)
// 	if err != nil {
// 		return fmt.Errorf("failed to write config file: %w", err)
// 	}

// 	return nil
// }

// func mapToStruct(configMap map[string]interface{}) *Cfg {
// 	config := &Cfg{}
// 	for key, value := range configMap {
// 		switch key {
// 		case "EnableTLS":
// 			config.EnableTLS, _ = strconv.ParseBool(value.(string))
// 		case "EnableNoTLS":
// 			config.EnableNoTLS, _ = strconv.ParseBool(value.(string))
// 		case "MaxUploadSize":
// 			config.MaxUploadSize, _ = strconv.ParseInt(value.(string), 10, 64)
// 		case "DaysOld":
// 			config.DaysOld, _ = strconv.Atoi(value.(string))
// 		case "ServerPortTLS":
// 			config.ServerPortTLS = value.(string)
// 		case "ServerPort":
// 			config.ServerPort = value.(string)
// 		case "CertPathCrt":
// 			config.CertPathCrt = value.(string)
// 		case "CertPathKey":
// 			config.CertPathKey = value.(string)
// 		case "BindtoAdress":
// 			config.BindtoAdress = value.(string)
// 		case "MaxVideosPerHour":
// 			config.MaxVideosPerHour, _ = strconv.Atoi(value.(string))
// 		case "MaxVideoNameLen":
// 			config.MaxVideoNameLen, _ = strconv.Atoi(value.(string))
// 		case "VideoResLow":
// 			config.VideoResLow = value.(string)
// 		case "VideoResMed":
// 			config.VideoResMed = value.(string)
// 		case "VideoResHigh":
// 			config.VideoResHigh = value.(string)
// 		case "BitRateLow":
// 			config.BitRateLow = value.(string)
// 		case "BitRateMed":
// 			config.BitRateMed = value.(string)
// 		case "BitRateHigh":
// 			config.BitRateHigh = value.(string)
// 		case "EnableFDP":
// 			config.EnableFDP, _ = strconv.ParseBool(value.(string))
// 		case "EnablePHL":
// 			config.EnablePHL, _ = strconv.ParseBool(value.(string))
// 		case "UploadPath":
// 			config.UploadPath = value.(string)
// 		case "ConvertPath":
// 			config.ConvertPath = value.(string)
// 		case "CheckOldEvery":
// 			config.CheckOldEvery = value.(string)
// 		case "AllowUploadOnlyFromUsers":
// 			config.AllowUploadOnlyFromUsers, _ = strconv.ParseBool(value.(string))
// 		case "AllowUploadOnlyFromAdmins":
// 			config.AllowUploadOnlyFromAdmins, _ = strconv.ParseBool(value.(string))
// 		case "VideoOnlyForUsers":
// 			config.VideoOnlyForUsers, _ = strconv.ParseBool(value.(string))
// 		case "NrOfCoreVideoConv":
// 			config.NrOfCoreVideoConv = value.(string)
// 		case "DelVidAftUpl":
// 			config.DelVidAftUpl, _ = strconv.ParseBool(value.(string))
// 		case "VideoPerPage":
// 			config.VideoPerPage, _ = strconv.Atoi(value.(string))
// 		case "VideoConvPreset":
// 			config.VideoConvPreset = value.(string)
// 		case "AllowEmbedded":
// 			config.AllowEmbedded, _ = strconv.ParseBool(value.(string))
// 		}
// 	}
// 	return config
// }

// func deleteOLD() {
// 	// Create a ticker to check for old files every `checkOldEvery` seconds
// 	ticker := time.NewTicker(checkOldEvery)

// 	// Start a goroutine to handle the ticker events
// 	go func() {
// 		for range ticker.C {
// 			// Delete old files in the upload path
// 			go deleteOldFiles(AppConfig.UploadPath, AppConfig.DaysOld)

// 			// Delete old files in the convert path
// 			go deleteOldFiles(AppConfig.ConvertPath, AppConfig.DaysOld)
// 		}
// 	}()

// 	// Wait for the goroutine to finish
// 	ticker.Stop()
// }

// // deleteOldFiles removes files and folders within the given folderPath that are older than the specified daysOld.
// func deleteOldFiles(folderPath string, daysOld int) error {
// 	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}

// 		if info.IsDir() {
// 			if time.Since(info.ModTime()).Hours()/24 >= float64(daysOld) {
// 				if err := os.RemoveAll(path); err != nil {
// 					return err
// 				}
// 				fmt.Printf("Folder %q deleted.\n", path)
// 				return filepath.SkipDir
// 			}
// 			return nil
// 		}

// 		if time.Since(info.ModTime()).Hours()/24 >= float64(daysOld) {
// 			if err := os.Remove(path); err != nil {
// 				return err
// 			}
// 			fmt.Printf("File %q deleted in folder %q.\n", info.Name(), folderPath)
// 		}
// 		return nil
// 	})
// 	return err
// }

// func resetVideoUploadedCounter() {
// 	// Create an atomic integer to store the counter
// 	var videosUploaded atomic.Int64

// 	// Start a goroutine to reset the counter every hour
// 	go func() {
// 		for range time.NewTicker(time.Hour).C {
// 			videosUploaded.Store(0)
// 		}
// 	}()

// 	// Wait for the goroutine to finish
// 	time.Sleep(time.Hour)
// }
