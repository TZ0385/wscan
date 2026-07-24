/**
* @Author: shaochuyu
* @Date: 5/7/2022 11:30
 */
package gopoc

import (
	"context"
	"wscan/core/model"
	logger "wscan/core/utils/log"

	"wscan/core/plugins/base"
)

type tomcatPut struct{}

// poc-go-tomcat-put
// id __golang gunkit_core_assassin_plugins_phantasm_pocs_gopoc___ptr_tomcatPut__Finger_func2(__int128 a1, __int64 a2)
func (*tomcatPut) Finger() *base.Finger {
	return &base.Finger{
		ExecAction: func(ctx context.Context, ab *base.Apollo) error {
			flow := ab.GetTargetFlow()
			logger.Infof("Start detection [%s] URL=%s", "poc-go-tomcat-put", flow.Request.URL().String())

			return nil
		},
		Channel: "web-directory",
		Binding: &model.VulnBinding{ID: "poc-go-tomcat-put", Plugin: "poc-go-tomcat-put", Category: "poc"},
	}

	//  *(_QWORD *)(v3 + 16) = "websitewedbar;wedgeq;weierp;";

	//gunkit_core_assassin_plugins_base___ptr_Bifrost__GetTargetFlow((_QWORD *)a2);
	//  v52 = *v11;
	//  gunkit_core_assassin_http___ptr_Request__clone(v52, 1);
	//  v54 = v17;
	//  gunkit_core_assassin_utils_RandLowerLetter(6LL, (_DWORD)v11);
	//  v55 = v11;
	//  v50 = github_com_thoas_go_funk_RandomString(12LL, 0LL);
	//  v56 = github_com_thoas_go_funk_RandomString(12LL, 0LL);
	//  runtime_convTstring(v50);
	//  v68[0] = &unk_1539F60;
	//  v68[1] = 0LL;
	//  runtime_convTstring(v56);
	//  v68[2] = &unk_1539F60;
	//  v68[3] = 0LL;
	//  fmt_Sprintf((__int64)"%s <%%-- %s --%%>", 17LL, (__int64)v68, 2LL, 2LL);
	//  v58 = v32;
	//  runtime_newobject(&unk_17997A0);
	//  v2 = *(_QWORD *)(v52 + 80);
	//  v62 = (_QWORD *)v12;
	//  if ( dword_3197BF0 )
	//  {
	//    v18 = v2;
	//    runtime_typedmemmove((__int64)&unk_17997A0);
	//  }
	//  else
	//  {
	//    ((void (*)(void))loc_472522)();
	//  }
	//  v65 = 0LL;
	//  runtime_convTstring((__int64)v55);
	//  *(_QWORD *)&v65 = &unk_1539F60;
	//  *((_QWORD *)&v65 + 1) = v18;
	//  v45 = fmt_Sprintf((__int64)"/%s.jsp/", 8LL, (__int64)&v65, 1LL, 1LL);
	//  v3 = v62;
	//  v62[8] = v40;
	//  v4 = (__int64)(v3 + 7);
	//  if ( dword_3197BF0 )
	//    runtime_gcWriteBarrier();
	//  else
	//    v3[7] = v32;
	//  v3[13] = 0LL;
	//  if ( dword_3197BF0 )
	//  {
	//    runtime_gcWriteBarrierBX();
	//    v4 = v10;
	//  }
	//  else
	//  {
	//    v3[12] = 0LL;
	//  }
	//  v61 = v4;
	//  v3[15] = 0LL;
	//  if ( dword_3197BF0 )
	//    runtime_gcWriteBarrierBX();
	//  else
	//    v3[14] = 0LL;
	//  v33 = gunkit_core_assassin_http___ptr_Request__WithURL(v52, (__int64)v3, v19);
	//  v51 = v20;
	//  v28 = runtime_stringtoslicebyte(0LL, v58);
	//  v60 = v25;
	//  runtime_newobject("(");
	//  v13[1] = v28;
	//  v13[2] = v33;
	//  if ( dword_3197BF0 )
	//    runtime_gcWriteBarrier();
	//  else
	//    *v13 = v60;
	//  v13[3] = 0LL;
	//  v13[4] = -1LL;
	//  gunkit_core_assassin_http___ptr_Request__WithBody(
	//    v51,
	//    (__int64)&off_2495A60,
	//    (__int64)v13,
	//    (__int64)"application/octet-stream",
	//    24LL,
	//    v33);
	//  *(_QWORD *)(v34 + 8) = 3LL;
	//  if ( dword_3197BF0 )
	//    runtime_gcWriteBarrier();
	//  else
	//    *(_QWORD *)v34 = "PUT";
	//  *(_BYTE *)(v34 + 64) = 0;
	//  gunkit_core_assassin_http___ptr_Client__respond(**(_QWORD **)(a2 + 8), a1, *((__int64 *)&a1 + 1), v34, 1, v34);
	//  if ( v40 )
	//  {
	//    v5 = *(_QWORD *)qword_3156A90;
	//    v64[0] = *(_QWORD *)(v40 + 8);
	//    v64[1] = v45;
	//    github_com_kataras_golog___ptr_Logger__Log(v5, 2, v64);
	//  }
	//  else if ( *(_QWORD *)(v35 + 24) == 201LL )
	//  {
	//    strings_TrimRight(v62[7], v62[8], (__int64)"/", 1LL);
	//    v6 = (__int64)v62;
	//    v62[8] = v35;
	//    if ( dword_3197BF0 )
	//      runtime_gcWriteBarrier();
	//    else
	//      *(_QWORD *)(v6 + 56) = v29;
	//    gunkit_core_assassin_http___ptr_Request__WithURL(v54, v6, v21);
	//    v53 = v22;
	//    v22[1] = 3LL;
	//    if ( dword_3197BF0 )
	//      runtime_gcWriteBarrier();
	//    else
	//      *v22 = "GETGIDGT;";
	//    v7 = *(_QWORD *)qword_3156A90;
	//    v67[0] = "\b";
	//    v67[1] = v62;
	//    github_com_kataras_golog___ptr_Logger__Logf(v7, 5, &aTlsDialwithdia[2120], 30LL, v67, 1LL, 1LL);
	//    gunkit_core_assassin_http___ptr_Client__respond(
	//      **(_QWORD **)(a2 + 8),
	//      a1,
	//      *((__int64 *)&a1 + 1),
	//      (__int64)v53,
	//      1,
	//      v36);
	//    if ( v41 )
	//    {
	//      v8 = *(_QWORD *)qword_3156A90;
	//      v66[0] = *(_QWORD *)(v41 + 8);
	//      v66[1] = v45;
	//      github_com_kataras_golog___ptr_Logger__Log(v8, 2, v66);
	//    }
	//    else
	//    {
	//      v57 = v37;
	//      RawBody = gunkit_core_assassin_http___ptr_Response__GetRawBody(v37, v14, v23);
	//      v46 = v24;
	//      v59 = v15;
	//      v47 = v26;
	//      v30 = runtime_stringtoslicebyte((__int64)v49, v50);
	//      bytes_Index(v59, v24, v26, v26, v30, v37, RawBody);
	//      if ( v43 == -1 )
	//      {
	//        v9 = 0;
	//      }
	//      else
	//      {
	//        v31 = runtime_stringtoslicebyte((__int64)v48, v56);
	//        bytes_Index(v59, v46, v47, v27, v31, v38, v43);
	//        v9 = v44 == -1;
	//      }
	//      if ( v9 )
	//      {
	//        runtime_newobject("0");
	//        v63 = v16;
	//        if ( dword_3197BF0 )
	//        {
	//          runtime_typedmemclr((__int64)"0", v16);
	//          runtime_gcWriteBarrier();
	//          runtime_gcWriteBarrierDX();
	//        }
	//        else
	//        {
	//          *(_OWORD *)v16 = 0LL;
	//          *(_OWORD *)(v16 + 16) = 0LL;
	//          *(_OWORD *)(v16 + 32) = 0LL;
	//          *(_QWORD *)v16 = v54;
	//          *(_QWORD *)(v16 + 8) = v57;
	//        }
	//        runtime_newobject("\b");
	//        if ( dword_3197BF0 )
	//          runtime_gcWriteBarrier();
	//        else
	//          *(_QWORD *)v16 = v63;
	//        gunkit_core_assassin_plugins_base___ptr_Bifrost__NewWebVulnFromFlow(a2, v16, 1LL, 1LL, 0LL, v38);
	//        gunkit_core_assassin_plugins_base___ptr_Bifrost__OutputVuln(a2, v39);
	//      }
	//    }
	//  }
	return nil
}
