import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';

const fetchData = async (url) => {
  const response = await fetch(url);
  return response.json();
};

const fetchRelatedData = async (url, key) => {
  const data = await fetchData(url);
  return { [key]: data };
};

const fetchApiData = async (apiItem) => {
  const apiSettingUrl = `http://localhost:8080/api-setting/fetch?api_id=${apiItem.api_id}`;
  const apiResultHistoryUrl = `http://localhost:8080/api-result-history/fetch?api_id=${apiItem.api_id}`;
  const apiSettingData = await fetchData(apiSettingUrl);
  const apiResultHistoryData = await fetchData(apiResultHistoryUrl);

  const fetchApiAtherData = async (apiSettingItem) => {
    const apiHeaderUrl = `http://localhost:8080/api-header-setting/fetch?api_setting_id=${apiSettingItem.api_setting_id}`;
    const apiParamUrl = `http://localhost:8080/api-param-setting/fetch?api_setting_id=${apiSettingItem.api_setting_id}`;
    const apiHeaderData = await fetchRelatedData(apiHeaderUrl, 'api_header');
    const apiParamData = await fetchRelatedData(apiParamUrl, 'api_param');

    return { ...apiSettingItem, ...apiHeaderData, ...apiParamData };
  };

  const apiSettingWithData = await Promise.all(
    apiSettingData.map(async (apiSettingItem) => fetchApiAtherData(apiSettingItem)),
  );

  return { ...apiItem, api_setting: apiSettingWithData, api_result_history: apiResultHistoryData };
};

export const getApi = createAsyncThunk('api/getApi', async () => {
  const apiUrl = 'http://localhost:8080/api/fetch';
  const apiData = await fetchData(apiUrl);

  const apiDataWithRelatedData = await Promise.all(
    apiData.map(async (apiItem) => fetchApiData(apiItem))
  );

  return apiDataWithRelatedData;
});

export const apiSlice = createSlice({
  name: 'api',
  initialState: {
    api: [],
    loading: false,
    error: null,
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(getApi.pending, (state) => {
        state.loading = true;
      })
      .addCase(getApi.fulfilled, (state, action) => {
        state.loading = false;
        state.api = action.payload;
        console.log(state.api)
      })
      .addCase(getApi.rejected, (state) => {
        state.loading = false;
        state.error = true;
      });
  },
});

export default apiSlice.reducer;