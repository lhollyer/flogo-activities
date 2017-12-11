package statistica_activity

//
//#include <string.h> 
//
//#define STRING_EQUAL(a,b) (_stricmp(a,b)==0)
//#define STRING_NOT_EQUAL(a,b) (_stricmp(a,b)!=0)
//
//void _BTrees_Prediction_T19_10_16(
//      const double * _CURRENT_WEIGHT__,
//      const double * _HOUROFDAY__,
//      const char * _DRINK_ID__,
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
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += 1.000000 * 6.40189129335304e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += 1.000000 * 1.91328669502073e+001;
//
//    }
//    else {
//    _PredictProb[0] += 1.000000 * 1.27425828135842e+001;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 7.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.94615025454218e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 7.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 9.74536773048490e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.44697444633378e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.46500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -8.71520263099512e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.46500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.93434711303827e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.56096526883802e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 6.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.67534080983610e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 6.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 8.52446889817342e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 4.38491658928147e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.50500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -7.37637838697396e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.50500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.61039115914012e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.66878187680482e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 7.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.89908760716368e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 7.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 8.12264151400565e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.56188222205717e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.56500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -6.18781509305446e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.56500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.44139862546825e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 5.55901782607971e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 6.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.27111571173838e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 6.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 5.38013026823461e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.66089683295897e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.41500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -6.60584581339009e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.41500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.59188632845477e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.35768181332775e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 5.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.01703642479084e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 5.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 4.59334499747613e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.88801285183171e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.62500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.29670936928867e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.62500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.01388528027995e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.57617655363626e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 8.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.99428609707361e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 8.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 5.70558552987552e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.34548415882580e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.66500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.63596575434973e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.66500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.85883267632371e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.07181937068893e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.36500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -5.27162060112376e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.36500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.81378258372086e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -5.43629578256605e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 5.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.69204244206876e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 5.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.19442781054279e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -3.06036366040107e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.54500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.18861841209834e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.54500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.95132192075201e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -3.47804422028453e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 8.97150371647189e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -9.01776597871163e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 8.80994454841373e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 5.35000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.34750134975209e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 5.35000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.38889532537709e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.35243486603418e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 8.19043284682410e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -8.58129509910669e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.14278788355039e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 8.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.51063596006061e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 8.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 5.00366266589833e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.97126929029031e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 8.50426970655287e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -8.64423105373274e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -5.65296968736821e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.32500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.61899440451335e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.32500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.55295133187903e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.12860534752457e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 8.35000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.44660164939448e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 8.35000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 4.65531039575942e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.85749547347369e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 8.27052779612963e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -8.54044373123445e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -9.44407267984390e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.70500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.74318227709752e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.70500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.37606584515670e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -3.91813988071667e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.32500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.14205821916447e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.32500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.40712693518322e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.33288341306155e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 7.93693175804738e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -8.15107848899778e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -5.53548192344730e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 9.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.31738093659638e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 9.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 4.94973709330748e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.00629166280091e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 7.71782175423447e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -7.60101413189683e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 8.22481867533952e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 4.55000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.30848293947113e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 4.55000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.47849314662802e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.23948814811450e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 7.40019542119538e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -7.28752069707756e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.73447864710498e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.27500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.74954354182336e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.27500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.29294287179868e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.56681500851900e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.32500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.20411623984467e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.32500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.27158095406158e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.07564350751740e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 7.27591659460477e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -7.58057210552999e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -8.17490739635374e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 9.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.27200588040721e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 9.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 4.37327366610156e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -3.68703402009349e-004;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 7.14788971511131e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -6.74264180241871e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.40724944051144e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 9.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.19650000867127e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 9.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 4.39241408079184e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 5.75953309169088e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 6.85313173167047e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -6.58761053231117e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.77525522743315e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.25500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.05898047794613e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.25500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.20752473054272e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.36733038466259e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 6.62566249544111e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -6.67662891415568e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 5.34490796351014e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.24500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.50472904577613e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.24500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.06295662252298e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -4.46733440962814e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 6.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.49742553879866e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 6.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.62095844728986e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.44093401613930e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 6.51996042991834e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -6.59558200121390e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -8.87649300997898e-004;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 9.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.07121097623165e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 9.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 4.48543408590190e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.97947760340610e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.69500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.21951265517364e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.69500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.89461869053131e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.17787954479774e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 6.56656913860186e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -6.58400954275695e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.26226727123882e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.01613340995467e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 4.25000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.12097810768425e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.06107051671665e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.20500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.81389567091349e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.20500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 9.65966340068065e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.12283377150905e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 6.82533139074441e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -6.29789687253441e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.04178234194420e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 9.55000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -9.86497259595188e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 9.55000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.64919805994824e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -3.57853633265844e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 6.32041043426635e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -6.09586649154227e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.39688233784362e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.25500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.89907514290982e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.25500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.01949956000142e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.51220616870371e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 5.76745528092100e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -5.98838336352253e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -8.91597014659864e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.00500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -9.51384346265361e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.00500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.87262238286851e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -8.24167679350360e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 5.83554249433562e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -5.91402012517852e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.68762769353198e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.19500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.68999918960120e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.19500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 9.12957768441916e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.16848026931345e-004;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 5.60849685654509e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -5.75407612424155e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -6.25593757733928e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 9.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -9.42810813529395e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 9.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.52005830451154e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.09171526776915e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.02500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -8.26921716365373e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.02500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.80792440941903e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.94797899517643e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.75500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -1.66254372869257e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.75500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.84654886684087e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 8.09463150099155e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 5.83929857501142e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -5.95134834547839e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.31367795046664e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.16500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.60947594123489e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.16500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 8.57520912545766e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -4.79056417214358e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 5.46251428119275e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -5.69852187182857e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -5.91427450251142e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.04500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -7.94159965359434e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.04500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.52490007221919e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 5.26074492106011e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 5.15381769401931e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -5.42095148621368e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.23324537993577e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.19500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.50278403764962e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.19500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 8.37095964278906e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.27088450371819e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 5.32236754463740e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.87157483236052e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.17384350078461e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.82457557374146e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 3.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.40142451783387e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -8.90812296081409e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.77500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -1.39667077143668e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.77500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.72370077708805e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 8.98582788678672e-003;
//
//    }
//    if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ <= 1.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.08007554810176e+000;
//
//    }
//    else if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ > 1.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.48757002603689e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 7.98246156064354e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.12500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.30647344411577e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.12500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 7.27274788628455e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -3.40199304759214e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 5.22159731780913e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -5.23073714096569e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 4.19893977776414e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.02500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -7.61447728134850e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.02500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.38599866367498e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.01864287346630e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 4.93423838933036e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -5.32509836435701e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.84833828162407e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.14500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.34073608728432e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.14500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 6.99030872488428e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.04259962221038e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 4.80399423435564e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -5.00175968917920e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -8.04023882326083e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.06500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -7.23759159766239e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.06500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.20001297375868e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -5.45471918182712e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 4.40321734351936e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.83235091578448e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.77977068167901e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.05500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -6.90493123666231e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.05500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.25054512783981e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.33836679733308e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 4.45567317010260e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.68097158343223e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -7.87333847385765e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.12500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.01525247430967e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.12500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 7.29008446537473e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.86304727669268e-002;
//
//    }
//    if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ <= 1.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.00150742929036e+000;
//
//    }
//    else if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ > 1.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -2.24451729657918e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.18363381230812e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.64924169589732e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 3.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.10317974195755e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.08606318509237e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 4.63782911528389e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.64040590499143e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 5.14012151687553e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.06500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -6.55195349125489e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.06500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.93880229863287e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.17042717333148e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 4.45617600231696e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.46798330361805e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.60706309441188e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.17500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.05852652157799e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.17500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 7.38500707508226e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -4.44500877256713e-004;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 4.43812913778394e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.32225020807110e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 7.06234117857599e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.12500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.92104736986841e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.12500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 6.89633037718307e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.44694275843581e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.82500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -1.10977045152382e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.82500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.66211756170018e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.49592139340585e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.06500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -6.31577083688716e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.06500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.69421906637510e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.29484962985620e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 4.41605928430216e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.71041991787898e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -8.72224066858943e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.09500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -5.83365685250176e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.09500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.69219408637929e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -6.43601833011652e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 4.40332959171783e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.35023315739327e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.66809831978344e-003;
//
//    }
//    if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ <= 1.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 3.51980705595918e-001;
//
//    }
//    else if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ > 1.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -5.14929634133908e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.91668209425003e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 4.04822082204084e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.93373960205517e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 5.65454383830013e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.08500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.99872532486370e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.08500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 6.10441207718442e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -8.08482418465275e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.13500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.71438679465080e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.13500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 6.00170466091767e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -4.16409321720347e-003;
//
//    }
//    if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ <= 1.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 8.80055331192908e-001;
//
//    }
//    else if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ > 1.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.71237812111633e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.70076122384173e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 4.00724627512327e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.24877544601283e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.16067505253328e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.11500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -5.80610339907953e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.11500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.93077174516322e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 7.02837400298502e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.72455534899393e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.30865059223538e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.88154192586821e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 4.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.16245257250607e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 4.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.27209096152274e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -5.58438056778966e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.75500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -1.09035803381080e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.75500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.27849037505846e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.06067369692308e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.79360034448185e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.14958419256232e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.78516549851769e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.05500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -5.81080454997155e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.05500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.79682600509538e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.25231299236775e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.11500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -5.04562426224112e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.11500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.91786384218680e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.04202638434761e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.99254303997752e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.11690489516760e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -5.14304442519255e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.08500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.56289728156150e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.08500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 5.64370044451896e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 9.45897367474042e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.81844331917173e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.84383182281771e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.74412122965685e-004;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.09500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.73553981948170e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.09500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 5.31816675797455e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.67319088613035e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.85271067917680e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.66795565947128e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.45400437476423e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 9.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -5.76502459555970e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 9.95000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.30674515058960e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -6.31354295483075e-003;
//
//    }
//    if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ <= 1.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.78654322254623e-001;
//
//    }
//    else if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ > 1.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -4.33650363482958e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -8.35249804060273e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.57781330954835e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.59346290211472e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.78225620288639e-004;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.07500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.55158174219691e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.07500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 5.22083600920485e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -4.98978163179375e-003;
//
//    }
//    if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ <= 1.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 7.18777232291532e-001;
//
//    }
//    else if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ > 1.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.64914729843744e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.51822874665513e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 3.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.38454481815331e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 3.45000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.06962941594461e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.13488727102523e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.54516607441451e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.46607803060674e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 4.93783258997064e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.82500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -8.41998475099558e-002;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.82500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.46578023014458e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 9.96487136702081e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.14500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.74268857019446e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.14500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.55468159846060e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 8.23207728590946e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.62326524859649e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.48027664068429e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.01722808269539e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.05500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.64067773241682e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.05500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 4.87774636680027e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.33323805860238e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.14500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.62158095223486e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.14500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.34882709153253e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -4.96555179415173e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.46443898802380e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.80243753980917e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.56004949459863e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.21540345082385e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.26225172274596e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 4.22167094751844e-007;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.05500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.41892194266874e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.05500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 4.54855190318434e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -8.71582215189492e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.33633573580809e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.33175796989648e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.48624678244886e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.14500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.92839120345500e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.14500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.41393616279133e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -9.50416229990451e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.04500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.31278675962075e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.04500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 4.50178928943881e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.31393716140533e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.26475673074332e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.39184587258632e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -6.38336616277939e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.07500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.16431154582170e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.07500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 4.60409006045378e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -9.35226301630168e-004;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.14871184952228e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.09504910901905e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 4.99727363174978e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.14500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.40580168582983e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.14500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.23823532083000e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.45128783599943e-003;
//
//    }
//    if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ <= 1.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 6.99071482087437e-001;
//
//    }
//    else if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ > 1.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.45997588507410e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.57750113776966e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.98681695895138e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.11709489525845e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -3.73267349635326e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.03500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.36344914046577e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.03500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 4.02413271089827e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.79798440897015e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 3.35000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.26982130526969e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 3.35000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 7.07829010985729e-002;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -3.68204977685828e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.15858668808308e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.92522937536190e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.49290413346719e-002;
//
//    }
//    if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ <= 1.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.49437470937453e-001;
//
//    }
//    else if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ > 1.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.73287010816114e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.28610950621314e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.18500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.15289797860463e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.18500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.53856785069006e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.66601858299413e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.04790846533438e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.04203071913263e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.77390493384819e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.78482276085419e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.98353543150615e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -8.85059677725431e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.02500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.32476632325457e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.02500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 4.33692074616876e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.12363883629940e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.13500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.09126473751517e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.13500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.01461635817429e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -3.83476092793084e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.90574734226294e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.97365512675758e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 5.44882700230716e-004;
//
//    }
//    if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ <= 9.50000000000000e+000 ) {
//        _PredictProb[0] += _LearningRate * 1.29678125350181e+000;
//
//    }
//    else if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ > 9.50000000000000e+000 ) {
//        _PredictProb[0] += _LearningRate * -7.00386042449403e-002;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -3.58891763564617e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.01500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.11447799429169e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.01500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 4.03949663786038e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 2.78123257214801e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.98504582274308e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.76459622614711e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.28107424260445e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.13500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.08322762004032e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.13500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.02979315595317e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.38788600632826e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.76375324157735e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.71497268708032e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 6.30852604119762e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.17500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.94216197401856e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.17500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.19157687201902e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 4.18325016555948e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.80500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -7.41788545126966e-002;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.80500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.02639034227594e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.92982355744758e-004;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.71781441453876e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.97398919619344e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.25340140291231e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.01500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -1.94805567273246e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.01500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 4.11133664408886e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.50172861564237e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.71250388005606e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.88834944128381e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -7.46846249835841e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.19500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.81000243684184e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.19500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.23313278977962e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 5.63584274259748e-003;
//
//    }
//    if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ <= 1.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.95821304237043e-001;
//
//    }
//    else if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ > 1.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.88836027680174e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 5.37665827820106e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.76343289578152e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.84297932701874e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -4.41082577703457e-004;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.01500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -1.97461803248044e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.01500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.85949857768060e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 7.55228360193820e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.63200022044780e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.70860462560031e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.80285041031729e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.22500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.76709506935468e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.22500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.26610138756879e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 4.74762641128865e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.42148503773838e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.66793836620921e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.04031035798462e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.01500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.12016030878928e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.01500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.49899599071067e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.54903713607303e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.21500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.37946456011931e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.21500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.12739647685683e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.10849949022222e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.75532617740187e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.49738031280371e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.48282682825310e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.88500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -6.29092039717948e-002;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.88500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.26385929968870e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.06834129362072e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.82500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -7.16364704198399e-002;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.82500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.03743042743790e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -5.98542821858546e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 3.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -1.14212800089852e+000;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 3.15000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 5.04645932822789e-002;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -6.39463901014332e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.59951503308395e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.63479887046787e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -3.56249598641509e-004;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.00500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -1.83980812633253e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.00500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.55068805403750e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 7.12989172350618e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.43053790196146e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.39247230453270e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.23336917076333e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.21500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.31071280998679e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.21500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.96609519142955e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.03502767097097e-003;
//
//    }
//    if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ <= 9.50000000000000e+000 ) {
//        _PredictProb[0] += _LearningRate * 1.18344864751985e+000;
//
//    }
//    else if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ > 9.50000000000000e+000 ) {
//        _PredictProb[0] += _LearningRate * -6.66912358981012e-002;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.05470650012123e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.30590759462522e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.44275436526926e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -6.11668725471204e-003;
//
//    }
//    if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ <= 1.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 2.16314124878123e-001;
//
//    }
//    else if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ > 1.75000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.01273563242447e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 9.52468600690943e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.99500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -1.81647045132174e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.99500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.08879885393312e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -7.22226941803885e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.34486831922735e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.38552095870403e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 4.34230425087166e-004;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.98500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -1.90205522033988e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.98500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.32771732029601e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -4.01779850679708e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.43084358703978e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.46396649477730e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -9.54016562852435e-004;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.20500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -3.37535853851722e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.20500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.97810779475179e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 5.19226376202635e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -8.05263770852925e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 4.05000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 6.26669969805123e-002;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -4.34097508566028e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.59631990242630e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.37603053308908e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.21678935941459e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 8.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -4.10576640501655e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 8.65000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.19348126126462e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.03782802552825e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.00500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -1.77320845498574e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.00500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.16744147620277e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.87655604899876e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.37751447475988e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.35615018924809e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 3.40232734404496e-003;
//
//    }
//    if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ <= 9.50000000000000e+000 ) {
//        _PredictProb[0] += _LearningRate * 1.02081272182365e+000;
//
//    }
//    else if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ > 9.50000000000000e+000 ) {
//        _PredictProb[0] += _LearningRate * -4.49160235624782e-002;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 5.81966531203589e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.95500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -1.81909713820472e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.95500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.91025832917558e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -8.13147425219071e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.35783642760482e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.40304390144493e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -9.95395470592111e-004;
//
//    }
//    if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ <= 1.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * 1.82056635801703e-001;
//
//    }
//    else if ( _HOUROFDAY__ != NULL && *_HOUROFDAY__ > 1.85000000000000e+001 ) {
//        _PredictProb[0] += _LearningRate * -3.17273612679632e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.72564019779749e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.22500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.98338978623659e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.22500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.85312305645610e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 6.45546564492906e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.12358921340583e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.46461018416168e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -1.49940549750456e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.23500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.89270031609302e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.23500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.64136881495316e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -5.26877754575484e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 2.91500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -4.50532544469177e-002;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 2.91500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.26297800359848e+000;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -8.03835182176261e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.07071806419328e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.60725083208928e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -2.54970888615324e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.21500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.94392355557239e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.21500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 1.82878433348071e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 5.61598196768150e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.94500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -1.62709998633339e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.94500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 3.10572288102428e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 1.31269395884885e-002;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.47825519350065e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.32317449069781e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * 9.19347639957374e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.13761115916189e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.59500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -2.33588482421490e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -8.41828258527806e-003;
//
//    }
//    if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ <= 1.97500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * -1.74772185737888e-001;
//
//    }
//    else if ( _CURRENT_WEIGHT__ != NULL && *_CURRENT_WEIGHT__ > 1.97500000000000e+002 ) {
//        _PredictProb[0] += _LearningRate * 2.94919696010486e-001;
//
//    }
//    else {
//    _PredictProb[0] += _LearningRate * -5.19725564972927e-003;
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

	var i1 = C.double(float64(context.GetInput("_CURRENT_WEIGHT__").(float64)))
	var i2 = C.double(float64(context.GetInput("_HOUROFDAY__").(float64)))
	var i3 = C.CString(context.GetInput("_DRINK_ID__").(string))

	var result C.double
	C._BTrees_Prediction_T19_10_16(&i1, &i2, i3, &result)

	context.SetOutput("result", float64(result))

	return true, nil
}

