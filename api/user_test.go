package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	mockdb "github.com/Qianjiachen55/pgK8/db/mock"
	db "github.com/Qianjiachen55/pgK8/db/sqlc"
	"github.com/Qianjiachen55/pgK8/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)


type eqCreateUserParamsMatcher struct {
	x db.CreateUserParams
	password string
}


func (e eqCreateUserParamsMatcher) Matches(xArg interface{}) bool {
	x, ok := xArg.(db.CreateUserParams)
	if !ok {
		return false
	}

	err := util.CheckPassword(e.password, x.HashedPassword)
	if err != nil{
		return false
	}
	e.x.HashedPassword = x.HashedPassword

	// In case, some value is nil
	if e.x == (db.CreateUserParams{}) || x == (db.CreateUserParams{}) {
		return reflect.DeepEqual(e.x, x)
	}

	// Check if types assignable and convert them to common type
	x1Val := reflect.ValueOf(e.x)
	x2Val := reflect.ValueOf(x)

	if x1Val.Type().AssignableTo(x2Val.Type()) {
		x1ValConverted := x1Val.Convert(x2Val.Type())
		return reflect.DeepEqual(x1ValConverted.Interface(), x2Val.Interface())
	}

	return false


}




func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("is equal to %v (%T)", e.x, e.x)
}

func EqCreateUserParams(x db.CreateUserParams,password string) gomock.Matcher{
	return eqCreateUserParamsMatcher{x,password}
}


func randomUser(t *testing.T) (user db.User, password string) {
	password = util.RandomString(6)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)
	user = db.User{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	return
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer,user db.User){
	data,err := ioutil.ReadAll(body)
	require.NoError(t, err)

	gotUser := db.User{}
	err = json.Unmarshal(data,&gotUser)
	require.NoError(t, err)

	require.Equal(t, user.Username,gotUser.Username)
	require.Equal(t, user.FullName,gotUser.FullName)
	require.Equal(t, user.Email,gotUser.Email)
	require.Empty(t, gotUser.HashedPassword)
}

func TestCreateUserAPI(t *testing.T) {
	user, password := randomUser(t)

	testCase := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username": user.Username,
				"password": password,
				"full_name":user.FullName,
				"email": user.Email,
			},
			
			
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateUserParams{
					Username:       user.Username,
					FullName:       user.FullName,
					Email:          user.Email,
				}	
				store.EXPECT().CreateUser(gomock.Any(),EqCreateUserParams(arg,password)).Times(1).Return(user,nil)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t,http.StatusOK,recoder.Code)
				requireBodyMatchUser(t,recoder.Body,user)
			},
		},
	}

	for _,tc:= range testCase{
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recoder := httptest.NewRecorder()

			// Marchal body data to JSON
			data,err := json.Marshal(tc.body)
			require.NoError(t, err)

			url :="/users"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recoder,request)
			tc.checkResponse(recoder)

		})
	}

}
