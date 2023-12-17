import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { getApi } from "./apiSlice";
import ApiListContents from './ApiListContents';

function ApiList() {
  const dispatch = useDispatch();
  const { api, loading, error } = useSelector((state) => state.api);

  useEffect(() => {
    dispatch(getApi());
  }, [dispatch]);

  return (
    <div className="ApiList">
      <h1>ApiList</h1>
      <div className="ApiListContainer">
        {loading && <p>Loading</p>}
        {error && <p>データ取得に失敗しました。</p>}
        {api && api.map((api) => (
          <ApiListContents key={api.api_id} apiId={api.api_id} />
        ))}
      </div>
    </div>
  );
}

export default ApiList;