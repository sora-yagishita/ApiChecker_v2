import { useSelector } from 'react-redux';

function getlastResultHistory(targetApi) {
  var lastResultHistory = null;
  if (targetApi.api_result_history && targetApi.api_result_history.length > 0) {
    lastResultHistory = targetApi.api_result_history.reduce((max, current) => {
      const currentDateTime = new Date(current.request_date_time);
      const maxDateTime = max ? new Date(max.request_date_time) : null;
      if (!maxDateTime || currentDateTime > maxDateTime) {
        return current;
      } else {
        return max;
      }
    }, null);
  }
  return lastResultHistory;
}

function ApiListContents({ apiId }) {
  const { api } = useSelector((state) => state.api);
  const targetApi = api.filter((apiItem) => apiItem.api_id === apiId)[0];
  const lastResultHistory = getlastResultHistory(targetApi);
  console.log(lastResultHistory)

  return (
    <div className="ApiContents">
      <div className="ApiName">
        <h2>
          {targetApi.api_name}
        </h2>
      </div>
      <div className="ApiDescription">
        <span>Discription</span><br/>
        <span>
          {targetApi.api_description}
        </span>
      </div>
      <div className="ApiSetting">
        <div className="LastDone">
          <span>LastDone : </span>
          <span>
            {lastResultHistory ?  lastResultHistory.request_date_time : "Not Executed"}
          </span>
        </div>
        <div className="ApiSettingInfo">
          <div className="LastStatus">
            <span>
              {lastResultHistory ? lastResultHistory.response_status_code : "---"}
            </span>
          </div>
          <div className="EndPoint">
            <span className="EndPointText">
              {targetApi.api_setting[0].endpoint}
            </span>
          </div>
          <div className="SrndBtn">
            <span>SEND</span>
          </div>
        </div>
      </div>
    </div>
  );
}

export default ApiListContents;