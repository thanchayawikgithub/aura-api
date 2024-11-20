package test

type TestCase[Args any, Want any] struct {
	Name    string
	Mock    func()
	Args    Args
	Want    Want
	WantErr bool
	Err     error
}

// func RunTestCase[Args any, Want any](
// 	suite *suite.Suite,
// 	testCase TestCase[Args, Want],
// 	testFunc func(Args) (Want, error),
// ) {
// 	suite.Run(testCase.Name, func() {
// 		if testCase.Mock != nil {
// 			testCase.Mock()
// 		}

// 		got, err := testFunc(testCase.Args)
// 		if testCase.WantErr {
// 			suite.Error(err)
// 			suite.Equal(testCase.Err, err)
// 		} else {
// 			suite.NoError(err)
// 			suite.Equal(testCase.Want, got)
// 		}
// 	})
// }

// func RunTestCase[Args any, Want any](
// 	suite *suite.Suite,
// 	testCase TestCase[Args, Want],
// 	testFunc interface{},
// ) {
// 	suite.Run(testCase.Name, func() {
// 		if testCase.Mock != nil {
// 			testCase.Mock()
// 		}

// 		fnValue := reflect.ValueOf(testFunc)
// 		args := []reflect.Value{
// 			reflect.ValueOf(testCase.Args),
// 		}

// 		results := fnValue.Call(args)

// 		got := results[0].Interface()
// 		err, _ := results[1].Interface().(error)

// 		if testCase.WantErr {
// 			suite.Error(err)
// 			suite.Equal(testCase.Err, err)
// 		} else {
// 			suite.NoError(err)
// 			suite.Equal(testCase.Want, got)
// 		}
// 	})
// }
