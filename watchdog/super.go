package watchdog

import (
	"errors"
	"strings"
)

func GetSuperStatus() (map[string]string, error) {
	args := []string{"status", "|", "awk", `'{print $1"="$2}'`}
	statusStrings, err := execCommand("supervisorctl", args)
	if err != nil {
		return nil, errors.New("读取失败" + err.Error())
	}
	//	statusStrings := `brm_consumer=RUNNING
	//castle_consumer:castle_consumer-0=RUNNING
	//castle_payment_notify_fms=RUNNING
	//castle_super_mass_message_scan=RUNNING
	//castle_super_mass_message_send=STARTING
	//fms_beanstalk_consumer=RUNNING
	//fms_consumer=RUNNING
	//ieas_arch_data_pass_summary=RUNNING
	//ieas_beanstalk_consumer=RUNNING
	//ieas_booth_buyer_summary=RUNNING
	//ieas_booth_data_pull=FATAL
	//ieas_booth_device_summary=RUNNING
	//ieas_booth_exhibitor_buyer_summary=RUNNING
	//ieas_consumer=RUNNING
	//ieas_gateway_worker=FATAL
	//`
	ss := strings.Split(statusStrings, "\n")
	sMap := make(map[string]string)
	for _, s := range ss {
		split := strings.Split(s, "=")
		if len(split) == 2 {
			sMap[split[0]] = split[1]
		}
	}
	return sMap, nil
}
