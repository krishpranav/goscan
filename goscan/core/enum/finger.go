package enum

import (
	"fmt"
	"github.com/krishpranav/goscan/core/utils"
)

func (s *EnumScan) EnumFINGER() {
	for _, port := range s.Target.GetPorts(utils.Config.DB) {
		// Enumerate only if port is open
		if port.Status == "open" {
			// Dispatch the correct scanner
			if port.Number == 79 {
				// Start Enumerating
				utils.Config.Log.LogInfo(fmt.Sprintf("Starting Enumeration: %s:%d (%s)", s.Target.Address, port.Number, "finger"))
				name := fmt.Sprintf("%s_finger_nmap_%d", s.Target.Address, port.Number)
				nmapArgs := fmt.Sprintf("-sV -Pn --script=finger -p%d", port.Number)
				s.runNmap(name, s.Target.Address, "FINGER", name, nmapArgs)
				output := s.makeOutputPath("FINGER", fmt.Sprintf("%s_finger_user-enum", s.Target.Address))
				cmd := fmt.Sprintf("finger-user-enum.pl -U %s -t %s > %s", utils.WORDLIST_FINGER_USER, s.Target.Address, output)
				s.runCmd(cmd)
			}
		}
	}
}