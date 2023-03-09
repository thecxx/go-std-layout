// Copyright 2023 Kami
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	"fmt"
	"strings"
)

func GenerateApacheLicense2(years string, owners ...string) (tpl string) {
	if len(owners) <= 0 {
		panic("owner no found")
	}
	tpl += "                                 Apache License\n"
	tpl += "                           Version 2.0, January 2004\n"
	tpl += "                        http://www.apache.org/licenses/\n"
	tpl += "\n"
	tpl += "   TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION\n"
	tpl += "\n"
	tpl += "   1. Definitions.\n"
	tpl += "\n"
	tpl += "      \"License\" shall mean the terms and conditions for use, reproduction,\n"
	tpl += "      and distribution as defined by Sections 1 through 9 of this document.\n"
	tpl += "\n"
	tpl += "      \"Licensor\" shall mean the copyright owner or entity authorized by\n"
	tpl += "      the copyright owner that is granting the License.\n"
	tpl += "\n"
	tpl += "      \"Legal Entity\" shall mean the union of the acting entity and all\n"
	tpl += "      other entities that control, are controlled by, or are under common\n"
	tpl += "      control with that entity. For the purposes of this definition,\n"
	tpl += "      \"control\" means (i) the power, direct or indirect, to cause the\n"
	tpl += "      direction or management of such entity, whether by contract or\n"
	tpl += "      otherwise, or (ii) ownership of fifty percent (50%) or more of the\n"
	tpl += "      outstanding shares, or (iii) beneficial ownership of such entity.\n"
	tpl += "\n"
	tpl += "      \"You\" (or \"Your\") shall mean an individual or Legal Entity\n"
	tpl += "      exercising permissions granted by this License.\n"
	tpl += "\n"
	tpl += "      \"Source\" form shall mean the preferred form for making modifications,\n"
	tpl += "      including but not limited to software source code, documentation\n"
	tpl += "      source, and configuration files.\n"
	tpl += "\n"
	tpl += "      \"Object\" form shall mean any form resulting from mechanical\n"
	tpl += "      transformation or translation of a Source form, including but\n"
	tpl += "      not limited to compiled object code, generated documentation,\n"
	tpl += "      and conversions to other media types.\n"
	tpl += "\n"
	tpl += "      \"Work\" shall mean the work of authorship, whether in Source or\n"
	tpl += "      Object form, made available under the License, as indicated by a\n"
	tpl += "      copyright notice that is included in or attached to the work\n"
	tpl += "      (an example is provided in the Appendix below).\n"
	tpl += "\n"
	tpl += "      \"Derivative Works\" shall mean any work, whether in Source or Object\n"
	tpl += "      form, that is based on (or derived from) the Work and for which the\n"
	tpl += "      editorial revisions, annotations, elaborations, or other modifications\n"
	tpl += "      represent, as a whole, an original work of authorship. For the purposes\n"
	tpl += "      of this License, Derivative Works shall not include works that remain\n"
	tpl += "      separable from, or merely link (or bind by name) to the interfaces of,\n"
	tpl += "      the Work and Derivative Works thereof.\n"
	tpl += "\n"
	tpl += "      \"Contribution\" shall mean any work of authorship, including\n"
	tpl += "      the original version of the Work and any modifications or additions\n"
	tpl += "      to that Work or Derivative Works thereof, that is intentionally\n"
	tpl += "      submitted to Licensor for inclusion in the Work by the copyright owner\n"
	tpl += "      or by an individual or Legal Entity authorized to submit on behalf of\n"
	tpl += "      the copyright owner. For the purposes of this definition, \"submitted\"\n"
	tpl += "      means any form of electronic, verbal, or written communication sent\n"
	tpl += "      to the Licensor or its representatives, including but not limited to\n"
	tpl += "      communication on electronic mailing lists, source code control systems,\n"
	tpl += "      and issue tracking systems that are managed by, or on behalf of, the\n"
	tpl += "      Licensor for the purpose of discussing and improving the Work, but\n"
	tpl += "      excluding communication that is conspicuously marked or otherwise\n"
	tpl += "      designated in writing by the copyright owner as \"Not a Contribution.\"\n"
	tpl += "\n"
	tpl += "      \"Contributor\" shall mean Licensor and any individual or Legal Entity\n"
	tpl += "      on behalf of whom a Contribution has been received by Licensor and\n"
	tpl += "      subsequently incorporated within the Work.\n"
	tpl += "\n"
	tpl += "   2. Grant of Copyright License. Subject to the terms and conditions of\n"
	tpl += "      this License, each Contributor hereby grants to You a perpetual,\n"
	tpl += "      worldwide, non-exclusive, no-charge, royalty-free, irrevocable\n"
	tpl += "      copyright license to reproduce, prepare Derivative Works of,\n"
	tpl += "      publicly display, publicly perform, sublicense, and distribute the\n"
	tpl += "      Work and such Derivative Works in Source or Object form.\n"
	tpl += "\n"
	tpl += "   3. Grant of Patent License. Subject to the terms and conditions of\n"
	tpl += "      this License, each Contributor hereby grants to You a perpetual,\n"
	tpl += "      worldwide, non-exclusive, no-charge, royalty-free, irrevocable\n"
	tpl += "      (except as stated in this section) patent license to make, have made,\n"
	tpl += "      use, offer to sell, sell, import, and otherwise transfer the Work,\n"
	tpl += "      where such license applies only to those patent claims licensable\n"
	tpl += "      by such Contributor that are necessarily infringed by their\n"
	tpl += "      Contribution(s) alone or by combination of their Contribution(s)\n"
	tpl += "      with the Work to which such Contribution(s) was submitted. If You\n"
	tpl += "      institute patent litigation against any entity (including a\n"
	tpl += "      cross-claim or counterclaim in a lawsuit) alleging that the Work\n"
	tpl += "      or a Contribution incorporated within the Work constitutes direct\n"
	tpl += "      or contributory patent infringement, then any patent licenses\n"
	tpl += "      granted to You under this License for that Work shall terminate\n"
	tpl += "      as of the date such litigation is filed.\n"
	tpl += "\n"
	tpl += "   4. Redistribution. You may reproduce and distribute copies of the\n"
	tpl += "      Work or Derivative Works thereof in any medium, with or without\n"
	tpl += "      modifications, and in Source or Object form, provided that You\n"
	tpl += "      meet the following conditions:\n"
	tpl += "\n"
	tpl += "      (a) You must give any other recipients of the Work or\n"
	tpl += "          Derivative Works a copy of this License; and\n"
	tpl += "\n"
	tpl += "      (b) You must cause any modified files to carry prominent notices\n"
	tpl += "          stating that You changed the files; and\n"
	tpl += "\n"
	tpl += "      (c) You must retain, in the Source form of any Derivative Works\n"
	tpl += "          that You distribute, all copyright, patent, trademark, and\n"
	tpl += "          attribution notices from the Source form of the Work,\n"
	tpl += "          excluding those notices that do not pertain to any part of\n"
	tpl += "          the Derivative Works; and\n"
	tpl += "\n"
	tpl += "      (d) If the Work includes a \"NOTICE\" text file as part of its\n"
	tpl += "          distribution, then any Derivative Works that You distribute must\n"
	tpl += "          include a readable copy of the attribution notices contained\n"
	tpl += "          within such NOTICE file, excluding those notices that do not\n"
	tpl += "          pertain to any part of the Derivative Works, in at least one\n"
	tpl += "          of the following places: within a NOTICE text file distributed\n"
	tpl += "          as part of the Derivative Works; within the Source form or\n"
	tpl += "          documentation, if provided along with the Derivative Works; or,\n"
	tpl += "          within a display generated by the Derivative Works, if and\n"
	tpl += "          wherever such third-party notices normally appear. The contents\n"
	tpl += "          of the NOTICE file are for informational purposes only and\n"
	tpl += "          do not modify the License. You may add Your own attribution\n"
	tpl += "          notices within Derivative Works that You distribute, alongside\n"
	tpl += "          or as an addendum to the NOTICE text from the Work, provided\n"
	tpl += "          that such additional attribution notices cannot be construed\n"
	tpl += "          as modifying the License.\n"
	tpl += "\n"
	tpl += "      You may add Your own copyright statement to Your modifications and\n"
	tpl += "      may provide additional or different license terms and conditions\n"
	tpl += "      for use, reproduction, or distribution of Your modifications, or\n"
	tpl += "      for any such Derivative Works as a whole, provided Your use,\n"
	tpl += "      reproduction, and distribution of the Work otherwise complies with\n"
	tpl += "      the conditions stated in this License.\n"
	tpl += "\n"
	tpl += "   5. Submission of Contributions. Unless You explicitly state otherwise,\n"
	tpl += "      any Contribution intentionally submitted for inclusion in the Work\n"
	tpl += "      by You to the Licensor shall be under the terms and conditions of\n"
	tpl += "      this License, without any additional terms or conditions.\n"
	tpl += "      Notwithstanding the above, nothing herein shall supersede or modify\n"
	tpl += "      the terms of any separate license agreement you may have executed\n"
	tpl += "      with Licensor regarding such Contributions.\n"
	tpl += "\n"
	tpl += "   6. Trademarks. This License does not grant permission to use the trade\n"
	tpl += "      names, trademarks, service marks, or product names of the Licensor,\n"
	tpl += "      except as required for reasonable and customary use in describing the\n"
	tpl += "      origin of the Work and reproducing the content of the NOTICE file.\n"
	tpl += "\n"
	tpl += "   7. Disclaimer of Warranty. Unless required by applicable law or\n"
	tpl += "      agreed to in writing, Licensor provides the Work (and each\n"
	tpl += "      Contributor provides its Contributions) on an \"AS IS\" BASIS,\n"
	tpl += "      WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or\n"
	tpl += "      implied, including, without limitation, any warranties or conditions\n"
	tpl += "      of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A\n"
	tpl += "      PARTICULAR PURPOSE. You are solely responsible for determining the\n"
	tpl += "      appropriateness of using or redistributing the Work and assume any\n"
	tpl += "      risks associated with Your exercise of permissions under this License.\n"
	tpl += "\n"
	tpl += "   8. Limitation of Liability. In no event and under no legal theory,\n"
	tpl += "      whether in tort (including negligence), contract, or otherwise,\n"
	tpl += "      unless required by applicable law (such as deliberate and grossly\n"
	tpl += "      negligent acts) or agreed to in writing, shall any Contributor be\n"
	tpl += "      liable to You for damages, including any direct, indirect, special,\n"
	tpl += "      incidental, or consequential damages of any character arising as a\n"
	tpl += "      result of this License or out of the use or inability to use the\n"
	tpl += "      Work (including but not limited to damages for loss of goodwill,\n"
	tpl += "      work stoppage, computer failure or malfunction, or any and all\n"
	tpl += "      other commercial damages or losses), even if such Contributor\n"
	tpl += "      has been advised of the possibility of such damages.\n"
	tpl += "\n"
	tpl += "   9. Accepting Warranty or Additional Liability. While redistributing\n"
	tpl += "      the Work or Derivative Works thereof, You may choose to offer,\n"
	tpl += "      and charge a fee for, acceptance of support, warranty, indemnity,\n"
	tpl += "      or other liability obligations and/or rights consistent with this\n"
	tpl += "      License. However, in accepting such obligations, You may act only\n"
	tpl += "      on Your own behalf and on Your sole responsibility, not on behalf\n"
	tpl += "      of any other Contributor, and only if You agree to indemnify,\n"
	tpl += "      defend, and hold each Contributor harmless for any liability\n"
	tpl += "      incurred by, or claims asserted against, such Contributor by reason\n"
	tpl += "      of your accepting any such warranty or additional liability.\n"
	tpl += "\n"
	tpl += "   END OF TERMS AND CONDITIONS\n"
	tpl += "\n"
	tpl += "   APPENDIX: How to apply the Apache License to your work.\n"
	tpl += "\n"
	tpl += "      To apply the Apache License to your work, attach the following\n"
	tpl += "      boilerplate notice, with the fields enclosed by brackets \"[]\"\n"
	tpl += "      replaced with your own identifying information. (Don't include\n"
	tpl += "      the brackets!)  The text should be enclosed in the appropriate\n"
	tpl += "      comment syntax for the file format. We also recommend that a\n"
	tpl += "      file or class name and description of purpose be included on the\n"
	tpl += "      same \"printed page\" as the copyright notice for easier\n"
	tpl += "      identification within third-party archives.\n"
	tpl += "\n"
	tpl += fmt.Sprintf("   Copyright %s %s\n", years, strings.Join(owners, ", "))
	tpl += "\n"
	tpl += "   Licensed under the Apache License, Version 2.0 (the \"License\");\n"
	tpl += "   you may not use this file except in compliance with the License.\n"
	tpl += "   You may obtain a copy of the License at\n"
	tpl += "\n"
	tpl += "       http://www.apache.org/licenses/LICENSE-2.0\n"
	tpl += "\n"
	tpl += "   Unless required by applicable law or agreed to in writing, software\n"
	tpl += "   distributed under the License is distributed on an \"AS IS\" BASIS,\n"
	tpl += "   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n"
	tpl += "   See the License for the specific language governing permissions and\n"
	tpl += "   limitations under the License.\n"
	return
}

func GenerateApacheLicense2Header(years string, owners ...string) (tpl string) {
	if len(owners) <= 0 {
		panic("owner no found")
	}
	tpl += fmt.Sprintf("Copyright %s %s\n", years, strings.Join(owners, ", "))
	tpl += "\n"
	tpl += "Licensed under the Apache License, Version 2.0 (the \"License\");\n"
	tpl += "you may not use this file except in compliance with the License.\n"
	tpl += "You may obtain a copy of the License at\n"
	tpl += "\n"
	tpl += "    http://www.apache.org/licenses/LICENSE-2.0\n"
	tpl += "\n"
	tpl += "Unless required by applicable law or agreed to in writing, software\n"
	tpl += "distributed under the License is distributed on an \"AS IS\" BASIS,\n"
	tpl += "WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n"
	tpl += "See the License for the specific language governing permissions and\n"
	tpl += "limitations under the License."
	return
}
