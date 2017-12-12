package stats

//
//#include <string.h>
//#include <math.h>
//
//#define STRING_EQUAL(a,b) (_stricmp(a,b)==0)
//#define STRING_NOT_EQUAL(a,b) (_stricmp(a,b)!=0)
//#define STRING_SET(a,b)   strcpy(a,b)
//
//
//void _GLZ_Prediction_T0_24_48(
//      const double * _CURRENT_WEIGHT__,
//      const double * _HOUROFDAY__,
//      double * _pRet
//   )
//   {
//    double _betas[3];
//    double _Covariates[2];
//    char* _DMatrix[2][3];
//
//    int _dumcode[2];
//    int _nparam;
//    int _npco;
//    double _val=0;
//    int _FactorID;
//    double _pval;
//    int _i=0, _j=0, _k=0, _ifind=0;
//    *_pRet=0;
//
//    for(_i=0; _i <2;_i++){
//      for(_j=0; _j <3;_j++){
//        _DMatrix[_i][_j] = "";
//      }
//    }
//
//    _betas[0] = 1.48341103823648e+000;
//    _betas[1] = 6.64013682413825e-003;
//    _betas[2] = -7.52260756804710e-003;
//
//    _Covariates[0] = *_CURRENT_WEIGHT__;
//    _Covariates[1] = *_HOUROFDAY__;
//    _DMatrix[0][1]="1";
//    _DMatrix[1][2]="1";
//
//    _dumcode[0]=1;
//    _dumcode[1]=-1;
//    _nparam = 3;
//    _npco=2;
//    if( _isnan(*_CURRENT_WEIGHT__) || _isnan(*_HOUROFDAY__) ){
//      *_pRet=0;
//      goto EndProgram;
//    }
//
//    for(_i = 0; _i <    _nparam;_i++){
//    _pval=1;
//    for(_j = 0; _j <_npco;_j++){
//        if(STRING_NOT_EQUAL(_DMatrix[_j][_i],"")){
//           if(_j <_npco){
//              _pval=_pval* (pow(_Covariates[_j] ,atof(_DMatrix[_j][_i])));
//           }
//        }
//     }
//     _val = _val + _betas[_i]*_pval;
//    }
//
//
//        *_pRet = exp(_val);
//
//    EndProgram:;
//
//}
//
import "C"

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log Is the Default package logger
var log = logger.GetLogger("activity-tibco-inference")

// InferfenceActivity is an Activity that is used to Invoke a ML Model using flogo-ml framework
type ModelActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a New InferfenceActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &ModelActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *ModelActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Runs an ML model
func (a *ModelActivity) Eval(context activity.Context) (done bool, err error) {

	var i1 = C.double(float64(context.GetInput("_CURRENT_WEIGHT__").(float64)))
	var i2 = C.double(float64(context.GetInput("_HOUROFDAY__").(float64)))

	var result C.double
	C._GLZ_Prediction_T0_24_48(&i1, &i2, &result)

	context.SetOutput("result", float64(result))

	return true, nil
}
