package statisticaactivity

//
//#include <string.h>
//
//#define STRING_EQUAL(a,b) (_stricmp(a,b)==0)
//#define STRING_NOT_EQUAL(a,b) (_stricmp(a,b)!=0)
//
//void _BTrees_Prediction_T15_26_44(
//      const double * _TEMPERATURE__,
//      const char * _EVENT__,
//      double * _pRet
//   )
//   {
//     double _LearningRate;
//     double _PredictProb[1];
//
//     *_pRet = 0;
//     _LearningRate = 0.100000;
//     _PredictProb[0] = 0;
//
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  ) {
//        _PredictProb[0] += 1.000000 * 5.04512889451211e+001;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  ) {
//        _PredictProb[0] += 1.000000 * 1.00000000000000e+002;
//
//    }
//    else {
//    _PredictProb[0] += 1.000000 * 5.10053170693563e+001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.44214415941673e+000;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -4.58527098104928e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 8.08827082919474e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.35000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -5.74498567966502e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.35000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.23379466868386e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.83780826538164e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.60164144635916e+000;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.04515456867106e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.10460702698210e+000;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.14548251980571e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 4.92505097704756e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 7.53324933625282e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -6.45654422062914e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 4.44243538360705e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 4.17582275826601e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 8.23875501028685e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.43364695258223e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.66786759467829e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.31935598497813e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -4.63642341039017e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.93511007647045e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.20396808078751e+000;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.94702208388467e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -3.59397127061404e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -9.83595535869818e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.46451507588816e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -8.85957322950060e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.55000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -8.88385650408355e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.55000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.41861537828908e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 8.09495238721030e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 9.80488221656005e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -4.10895287958531e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 4.44156100959941e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.45180209194943e+000;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.57960877360461e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -7.24472924552662e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -7.95785069463357e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.79247750719296e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -6.81565472826327e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 5.37558648231039e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -4.39553380882522e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.46323371171932e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.34324133492195e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -4.18843178302457e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -3.32809676886363e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 6.74719204109983e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -4.06856555350852e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.81953932470472e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 6.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 5.64273476134077e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 6.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 7.55117678213256e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 6.80431016266594e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  ) {
//        _PredictProb[0] += _LearningRate * -2.11694529733885e-001;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * 2.70067687827323e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.10801589297191e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.36799121113458e+000;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.16031429784758e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -4.28539996937467e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.55000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.62371828811999e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.55000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.26962623355182e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.07013405831677e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.79671132442920e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -4.00559035107583e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -7.63947417800258e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  ) {
//        _PredictProb[0] += _LearningRate * -1.03129758016093e+000;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * 2.33929703078822e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.89959414609519e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -4.39111521498107e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.97802708359592e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 4.82161602203495e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -6.85770522972021e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.86950146671730e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.63844334325239e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 6.48229779746197e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -4.88597842759318e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 7.04374164711347e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -7.29825248311524e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.99606028768491e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.43018968378629e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 6.10146139657241e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -4.67290222080256e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 7.57612275995186e-002;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  ) {
//        _PredictProb[0] += _LearningRate * -1.39966011525133e+000;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * 2.30827435747534e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -4.71829936493791e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.38938644459141e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.27753564785453e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.18445566888187e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 5.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 5.70377125004578e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 5.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 5.12633878035945e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 7.03536010559124e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 8.50781033551059e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.67215158305575e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.60995838929971e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 4.70397227990734e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.78679554799925e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.12816058359360e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  ) {
//        _PredictProb[0] += _LearningRate * -1.10075061811529e+000;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * 1.57313570292641e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -9.75411144114290e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 6.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.96049308717900e-002;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 6.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 6.20319214687170e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.26377490915715e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.50363501847858e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.95484777060131e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -5.80315512926201e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -9.28961568506863e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.14160483455276e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.99968891228130e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -5.31722401462947e-002;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.73755171154810e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 5.04885217257140e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 6.35000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.57476999225526e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 6.35000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 5.82843391352195e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.62125210819379e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 7.96795930254586e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -4.31570311554511e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.52054201093780e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 6.35000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.66247971014851e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 6.35000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 5.55096015395480e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.50092760596524e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 6.55000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.35837870717898e+000;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 6.55000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 5.18731644961957e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.25485861370344e+000;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 5.86484645727487e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.58941770410251e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.90807454681989e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.47223856228585e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.82450470372823e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -7.78022898415475e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -5.80838889212358e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.17294055145799e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.32791165711172e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.08016106365027e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.85215175229232e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.05606570985186e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 5.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -4.05696565845844e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 5.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 4.91147402225081e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.46875769485455e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -8.33478511815746e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.38357927527317e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -5.99469005219694e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 5.48873824535859e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.18292077911530e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.58471854421717e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.53663036025692e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.33297196300726e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 6.35426513403914e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.23637056876993e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.01705326578623e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.82645079247862e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.70083849315136e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.88002552413867e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.73061131293243e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 8.94278808802736e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.86392958083566e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 4.85030113005443e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.14180302688266e+000;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.07791398135088e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -6.83411659802270e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.11098049061236e-002;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.32488424170991e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 6.29234978404931e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 6.66004098695821e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.26913983306492e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.89974696005040e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 4.87109725979595e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.74049027747544e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 6.13998295866565e-003;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.55000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.69351375735060e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.55000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.91539982467074e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.43620828787363e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  ) {
//        _PredictProb[0] += _LearningRate * -4.84277512917831e-001;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * 3.70181119767557e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.40535354928654e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.23666149962512e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.80007676059571e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.52928910215886e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -5.85179276876290e-002;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.72124345581599e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.14235274293920e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 7.87145280357106e-002;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.22594767621160e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.64821561666419e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.40762485767635e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.11909789139671e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 8.77174163902066e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.96128275991298e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.62932640441046e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.67976583605861e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 5.23056932546579e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.17833569555878e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.35504914206161e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.55000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.06366293570696e+000;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.55000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.64327617697543e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.73563034363799e-003;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.55064314779129e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.98529133744134e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.06142124718263e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 5.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -7.21421628581159e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 5.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 4.88718592633722e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -5.54877823456915e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -4.73068436493142e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.72671124419711e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -6.68650031432328e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.19289979666315e+000;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.28544051452943e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -9.24576945201493e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 6.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.31455151026168e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 6.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 4.47712294228428e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.05127721093371e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.99350936527367e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.03960275638903e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.14948655737281e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  ) {
//        _PredictProb[0] += _LearningRate * -1.67591615862340e-001;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * 3.65683044778197e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.54392732348682e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 5.77537149614532e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.08630232228827e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 8.81976412943384e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.57505239399334e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.29436255756777e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.93770683804371e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -8.78376310581662e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.46863451431779e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -3.24780719507204e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -4.88336598904639e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.28601767034270e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -6.54458064307882e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.72822845210033e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.75353958894370e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 9.75140025096632e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.17106520573989e+000;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 6.73680950860803e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -7.89056856490429e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.07044507898869e+000;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.40465511683845e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 5.01082431961845e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  ) {
//        _PredictProb[0] += _LearningRate * 1.95258198797141e-001;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * 2.55439918910709e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.91058092750191e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.17015029666440e-002;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.03465958588883e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.88494418027716e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.92500098616907e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.92426008688201e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.94285516818810e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  ) {
//        _PredictProb[0] += _LearningRate * -1.00164811458734e+000;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * 2.25511387803662e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -7.39100486594915e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.41356871617966e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.49040735240970e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.00937397806806e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 9.41484551396447e-002;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.14868269121022e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 5.80181150810466e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -9.25742439655332e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.22585295262528e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.90781613403024e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.18987517661017e+000;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -5.62629836763962e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 9.27661163438534e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.33185732084717e+000;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.22759631237901e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 9.60176796398579e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  ) {
//        _PredictProb[0] += _LearningRate * -1.03384812678219e+000;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  ) {
//        _PredictProb[0] += _LearningRate * 1.22626963803417e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.98538656382953e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.42977529719280e-002;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.92006263835156e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -4.06272883678072e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  ) {
//        _PredictProb[0] += _LearningRate * -8.94548860175312e-001;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  ) {
//        _PredictProb[0] += _LearningRate * 8.96276439524607e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.34986088915216e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 4.12787825277067e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.49306050814470e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.97787325956930e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 5.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.95152496540655e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 5.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.65208977336894e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.75882095659269e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -5.24500306972594e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.53042171965664e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 4.58392130204372e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 5.70707486162723e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.28852407323169e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.77113918410084e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.47580825327575e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.37658587805326e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 7.57321081458257e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.35000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.86922585769834e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.35000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.48418644626661e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.63216436032575e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 8.80909524195521e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.65723483107329e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 5.04042655187126e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.21766312693437e+000;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 9.96735337537865e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -8.10573806197304e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.02125291578908e+000;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.32618047474085e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -7.68660246656265e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  ) {
//        _PredictProb[0] += _LearningRate * -5.19434560881684e-001;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * 2.94084371565831e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.20435147358333e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * -2.68297202117464e+000;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  ) {
//        _PredictProb[0] += _LearningRate * 7.54587688557605e-002;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.49954369777886e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  ) {
//        _PredictProb[0] += _LearningRate * -8.05804499294092e-001;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  ) {
//        _PredictProb[0] += _LearningRate * 3.10026526446575e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -9.83208862280905e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -9.53726761065061e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.54173542153996e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -4.45686217637805e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.11095036208151e+000;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.60302633792340e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -6.06787088296651e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 9.31220075019198e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -5.38599733487679e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 7.75217956870031e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 9.49364925731205e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.65415655036866e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 6.24958044677614e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 5.41237796788113e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.58002590551919e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.75124085191648e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 6.34363340585440e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.35599393863936e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.59618153893012e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.87676271656317e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.49074001536347e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.52661566304653e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -8.14313432059980e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.08666932514071e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -6.02566211295857e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 5.27436690011178e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.28963171363002e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.81150324504957e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.64879635576786e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.78816881604666e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 7.25137014094973e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 9.32200398350765e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.01101110747796e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 6.38952116562067e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.54783720605502e-002;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.00402659807604e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.91872213621971e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  ) {
//        _PredictProb[0] += _LearningRate * -1.54715261971576e-001;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * 3.14935551809432e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.17492149055923e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.95977580312893e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.26836039867234e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -9.11836149101754e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.61306690664301e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.86162088437807e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.02514299868389e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 5.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.98609645421145e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 5.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.60781404930796e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.83925532151115e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.13252168634750e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.78659405016633e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -5.76066503052775e-003;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.70004070496249e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.04496688255713e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.43031033159734e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 8.48880431629025e-002;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.40574421989860e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 6.41176687339895e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.35566129426551e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.95260288054970e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.81278387528962e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.93242073651717e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.23914638693749e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.38910674289566e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -7.64345268962372e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.65483090205694e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -5.83289837175638e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 8.88276799876663e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.06841943021670e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 4.33579640722611e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 5.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.32228182531142e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 5.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.98209485041018e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.17503681277865e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * -2.36004954603899e+000;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  ) {
//        _PredictProb[0] += _LearningRate * -9.44068294467630e-002;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.80865556397877e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.95743283238145e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.15181153322823e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 4.64185345292661e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.44277326356212e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.91472331588208e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -6.78067158911822e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 2.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.32211441869633e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 2.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.98044104830859e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -3.96471335762730e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -8.79877816555075e-002;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.51705943934479e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.46923646348474e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.13963486807591e+000;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 8.35152809072787e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -9.93723025526454e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.34144516784798e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.36143479677921e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -5.41477280299667e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.52673859011344e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.00528541858114e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.18160884856980e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.26511292828221e+000;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.40190275505362e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 8.65352592450463e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  ) {
//        _PredictProb[0] += _LearningRate * -1.65818021100147e+000;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  ) {
//        _PredictProb[0] += _LearningRate * 7.83711759256145e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.06246185550901e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  ) {
//        _PredictProb[0] += _LearningRate * 5.01498845545140e-001;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  ) {
//        _PredictProb[0] += _LearningRate * 1.40356642428479e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.06755714234577e+000;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.80791663671484e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.73516375675422e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.17949170188119e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  ) {
//        _PredictProb[0] += _LearningRate * -8.31795026358518e-001;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * 1.28196941127923e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -6.33375572044536e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  ) {
//        _PredictProb[0] += _LearningRate * 5.56025663009532e-001;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * 1.41057950312732e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.15962125021945e+000;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * -1.57146116579681e+000;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  ) {
//        _PredictProb[0] += _LearningRate * 6.60122913611430e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 4.74420608087834e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  ) {
//        _PredictProb[0] += _LearningRate * -5.61905354771034e-001;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * 1.70095740804459e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.70679117662798e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.02877475956499e+000;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 8.81960098780318e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -6.37559590301109e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.12900963162858e+000;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.75398813528129e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -8.48936714471509e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.53410390855737e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.87987303184316e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.48089507360275e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  ) {
//        _PredictProb[0] += _LearningRate * -1.74356860100823e+000;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  ) {
//        _PredictProb[0] += _LearningRate * 1.21491172413605e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.30577620534790e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  ) {
//        _PredictProb[0] += _LearningRate * 5.47286083426228e-001;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * 2.81956181280310e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 7.39291872799816e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -4.70152654177859e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.99869779991703e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.50234047523218e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -6.57857182285295e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 8.56559504437300e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -3.50531146382553e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.11712213041017e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.98936624702833e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.02323190879015e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.42683562367853e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 6.44446231330121e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.10509020842502e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.22357605814877e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.37210674912715e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.51151382993037e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.57453868959994e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.99150405087794e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -5.97300370898626e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.83551003833100e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.10392216212476e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -4.78485058182077e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -7.39491336682809e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.11339121953091e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -5.30967570516162e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  ) {
//        _PredictProb[0] += _LearningRate * -1.35901993901533e-001;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * 1.93639017120414e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.65171511328513e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -4.48103735045078e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 8.14621144227331e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.89528948365649e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -7.23048070071611e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.04684478541026e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -4.02686645830666e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.32579078482465e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.10699675794817e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 5.80582735757321e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.35000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.88016922517882e-002;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.35000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 8.63373352474604e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.87203594019593e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 6.51523561518295e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.14286213844717e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.78432791948611e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.28278230380774e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.35191252434866e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.23839210267535e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -7.88616826346468e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.05773373529790e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.13402474090556e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 6.60046591852774e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.69850554185094e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.15134135349831e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -8.11784920805825e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 8.14614217395015e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -5.10925271238834e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  ) {
//        _PredictProb[0] += _LearningRate * -4.76066832788700e-001;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  ) {
//        _PredictProb[0] += _LearningRate * 7.79562126894957e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.19966348564574e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -9.86460190894028e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.73887742751755e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -4.91527083057388e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.65810157101507e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.91317132948648e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 4.13042168485317e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 7.01673489526607e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.26676353427286e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 4.02521252156619e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.73513077081423e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.13300656990195e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.25969872906602e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -5.24591212146712e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.64974699727573e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.84034713971950e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -4.56379851432965e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.50547455256365e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.90011648510036e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * -1.35151105354775e+000;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  ) {
//        _PredictProb[0] += _LearningRate * 1.13374197061605e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 9.32783373121248e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * -2.06657333001710e+000;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  ) {
//        _PredictProb[0] += _LearningRate * 5.68060029259279e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.53531497888759e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  ) {
//        _PredictProb[0] += _LearningRate * -7.87314763916323e-001;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * 1.18092404161088e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -6.05237043780019e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 8.84249689661726e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.61331229194987e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 6.19965715157264e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.41560980823419e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.40720294934197e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 6.65192440144735e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 5.97642053568091e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.41248342381599e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.19541503986324e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.60725903628175e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -6.13645044206459e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.15520418715666e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * -2.24385093931369e+000;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  ) {
//        _PredictProb[0] += _LearningRate * 8.97220015702109e-002;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.02988815569681e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  ) {
//        _PredictProb[0] += _LearningRate * -1.45384623636937e-001;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * 1.03683675601501e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.92519826172909e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  ) {
//        _PredictProb[0] += _LearningRate * 1.52142519364409e-001;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * 4.80213644750527e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 5.28853657923711e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.44997635032524e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.56522478831420e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -3.01433560001333e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.35000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -9.58945858692281e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.35000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.08917881205536e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -4.47837125160379e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -4.66671370502606e-002;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.42575294028257e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -4.03935403351323e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 6.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 6.35381000458268e-002;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 6.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.80614492874558e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.08170145135528e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  ) {
//        _PredictProb[0] += _LearningRate * -3.97271184453212e-001;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * 3.49382403501151e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.21512442176684e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.35000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -5.75460615587366e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.35000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 4.20217406088756e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -3.22301037330298e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.11957678305781e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.35504215878938e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 6.21940393781121e-002;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.84160051432959e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.32934365587692e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.00869327891560e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -6.33909858033710e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.76767066636123e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -4.55267429583307e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.30705196772851e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.24157448625590e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 7.16933578521750e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -6.79178478744106e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 5.98481635238599e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -5.21060219495089e-001;
//
//    }
//    if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"PULL")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"INUSE")  || _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"TEMP")  ) {
//        _PredictProb[0] += _LearningRate * 2.26963476399987e-001;
//
//    }
//    else if ( _EVENT__ != NULL &&  STRING_EQUAL(_EVENT__,"VIBRATION")  ) {
//        _PredictProb[0] += _LearningRate * 2.95033304373833e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 4.53053726127885e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.08940134761950e+000;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.16184610865568e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -9.24952540989813e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 7.57513528592475e-001;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.74820581546138e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.94973816200529e-001;
//
//    }
//    if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ <= 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.79948055016832e-002;
//
//    }
//    else if ( _TEMPERATURE__ != NULL && *_TEMPERATURE__ > 3.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.63957954277025e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.69663789131782e-001;
//
//    }
//
//   *_pRet = _PredictProb[0];
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

	var i1 = C.double(float64(context.GetInput("_TEMPERATURE__").(float64)))
	var i2 = C.CString(context.GetInput("_EVENT__").(string))

	var result C.double
	C._BTrees_Prediction_T15_26_44(&i1, i2, &result)

	context.SetOutput("result", float64(result))

	return true, nil
}
