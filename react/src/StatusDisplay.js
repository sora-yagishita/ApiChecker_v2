import React, { useState, useEffect } from 'react';
import axios from 'axios';
import ApiResultGraph from './ApiResultGraph';

function StatusDisplay() {
  const [apiInfo, setApiInfo] = useState([]);
  const [statusCodes, setStatusCodes] = useState({});
  const [lastApiRequestTime, setLastApiRequestTime] = useState({});

  useEffect(() => {
    const endpointData = require('./endpoints.json');
    const apiInfo = endpointData.apiInfo;

    setApiInfo(apiInfo);

    apiInfo.forEach((api) => {
      fetchStatusCode(api);
      setInterval(() => fetchStatusCode(api), api.requestIntervalMin * 60000);
    });

    return () => {
      apiInfo.forEach((api) => {
        clearInterval(api.intervalId);
      });
    }
  }, []);

  const fetchStatusCode = async (api) => {
    const params = convertJson(api);
    axios.defaults.headers.common['Content-Type'] = 'application/json'

    await axios.get(
      api.apiEndpoint,
      {params}
    ).then(response => {
      setStatus(api.apiName, response.status);
      addAipRsult(api.apiName, response.status);
    }).catch(error => {
      setStatus(api.apiName, error.response.status);
      addAipRsult(api.apiName, error.response.status);
    })
  };

  const convertJson = (api) => {
    const apiParams = {};
    api.apiParams.forEach((param) => {
      apiParams[param.key] = param.value;
    });
    return apiParams;
  }

  const setStatus = (apiName, status) => {
    setStatusCodes((prevStatusCodes) => ({
      ...prevStatusCodes,
      [apiName]: status,
    }));
    setLastApiRequestTime((lastApiRequestTime) => ({
      ...lastApiRequestTime,
      [apiName]: new Date(),
    }));
  }

  const addAipRsult = async (apiName, status) => {
    await axios.post('http://localhost:8080/api-result/add-apiResult',
      {
        apiName: apiName,
        ApiStatus: String(status),
        ApiDateTime: new Date(),
      }
    ).then(response => {
    }).catch(error => {
      console.log(error);
    });
  }

  return (
    <div>
      {apiInfo.map((api) => (
        <div key={api.apiName}>
          <p>API名: {api.apiName}</p>
          <p>ステータスコード: {statusCodes[api.apiName]}</p>
          <p>最終更新時間: {lastApiRequestTime[api.apiName] ? lastApiRequestTime[api.apiName].toLocaleTimeString() : 'まだ送信されていません'}</p>
          <ApiResultGraph apiName={api.apiName} />
        </div>
      ))}
    </div>
  );
}

export default StatusDisplay;
