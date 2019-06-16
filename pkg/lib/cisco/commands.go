package cisco

func getPagerDisableCmd(device_type string) string {
	switch device_type {
	case "cisco_asa":
		return "terminal pager disable"

	default:
		return "terminal len 0"
	}

}

func getTermWidthCmd(device_type string) string {
	switch device_type {
	default:
		return "terminal width 511"

	}

}
