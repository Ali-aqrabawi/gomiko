package cisco

func getPagerDisableCmd(device_type string) string {
	switch device_type {
	case "cisco_asa":
		return "terminal pager 0"

	default:
		return "terminal len 0"
	}

}
