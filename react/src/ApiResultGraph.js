import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Line } from 'react-chartjs-2';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
} from 'chart.js';

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
);

export const options = {
  responsive: true,
  plugins: {
    legend: {
      position: 'top',
    },
    title: {
      display: true,
      text: 'Chart.js Line Chart',
    },
  },
};

function ApiResultGraph({ apiName }) {
  const [apiResults, setApiResults] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      const response = await axios.get('http://localhost:8080/api-result/fetch-apiResult', {
        params: {
          apiName: apiName,
        },
      }).then(response => {
        setApiResults(response.data);
      }).catch(error => {
        console.log(error);
      });
    };
    fetchData();
  }, [apiName]);


  const graphData = {
    labels: apiResults.map(item => item.apiDateTime),
    datasets: [
      {
        label: "ステータスコード",
        data: apiResults.map(item => Number(item.apiStatus)),
        borderColor: 'rgb(53, 162, 235)',
        backgroundColor: 'rgba(53, 162, 235, 0.5)',
      },
    ],
  };

  return (
    <div class="chart">
      <Line
        options={options}
        data={graphData}
      />
    </div>
  );
}

export default ApiResultGraph;