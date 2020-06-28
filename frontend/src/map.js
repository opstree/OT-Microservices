import React, { useState, useEffect } from 'react';
import { ComposableMap, Geographies, Geography } from 'react-simple-maps';
import { scaleQuantile } from 'd3-scale';
import ReactTooltip from 'react-tooltip';

import LinearGradient from './LinearGradient.js';
import './map.css';

/**
* Courtesy: https://rawgit.com/Anujarya300/bubble_maps/master/data/geography-data/india.topo.json
* Looking topojson for other countries/world? 
* Visit: https://github.com/markmarkoh/datamaps
*/
const INDIA_TOPO_JSON = require('./india.topo.json');

const PROJECTION_CONFIG = {
  scale: 350,
  center: [78.9629, 22.5937] // always in [East Latitude, North Longitude]
};

// Red Variants
const COLOR_RANGE = [
  '#ffedea',
  '#ffcec5',
  '#ffad9f',
  '#ff8a75',
  '#ff5533',
  '#e2492d',
  '#be3d26',
  '#8C92AC',
  '#92A1CF'
];

const DEFAULT_COLOR = '#EEE';

const geographyStyle = {
  default: {
    outline: 'none'
  },
  hover: {
    fill: '#ccc',
    transition: 'all 250ms',
    outline: 'none'
  },
  pressed: {
    outline: 'none'
  }
};

function generateNumber(input) {
  if (input === undefined) {
    return 0
  } else {
    return input
  }
}

// will generate random heatmap data on every call
function Map() {
  const [stats, handleStats] = useState([]);

  const FetchData = async () => {
    const jsondata = await fetch('http://172.17.0.3:8080/search/city');
    const stats = await jsondata.json();
    handleStats(stats)  
  }
  const newjsondata = stats
  console.log(newjsondata)

  const getHeatMapData = () => {
    return [
      { id: 'AP', state: 'Andhra Pradesh', value: generateNumber(newjsondata["Andhra Pradesh"]) },
      { id: 'AR', state: 'Arunachal Pradesh', value: generateNumber(newjsondata["Arunachal Pradesh"]) },
      { id: 'AS', state: 'Assam', value: generateNumber(newjsondata["Assam"]) },
      { id: 'BR', state: 'Bihar', value: generateNumber(newjsondata["Bihar"]) },
      { id: 'CT', state: 'Chhattisgarh', value: generateNumber(newjsondata["Chhattisgarh"]) },
      { id: 'GA', state: 'Goa', value: generateNumber(newjsondata["Goa"]) },
      { id: 'GJ', state: 'Gujarat', value: generateNumber(newjsondata["Gujarat"]) },
      { id: 'HR', state: 'Haryana', value: generateNumber(newjsondata["Haryana"]) },
      { id: 'HP', state: 'Himachal Pradesh', value: generateNumber(newjsondata["Himachal Pradesh"]) },
      { id: 'JH', state: 'Jharkhand', value: generateNumber(newjsondata["Jharkhand"]) },
      { id: 'KA', state: 'Karnataka', value: generateNumber(newjsondata["Karnataka"]) },
      { id: 'KL', state: 'Kerala', value: generateNumber(newjsondata["Kerala"]) },
      { id: 'MP', state: 'Madhya Pradesh', value: generateNumber(newjsondata["Madhya Pradesh"]) },
      { id: 'MH', state: 'Maharashtra', value: generateNumber(newjsondata["Maharashtra"]) },
      { id: 'MN', state: 'Manipur', value: generateNumber(newjsondata["Manipur"]) },
      { id: 'ML', state: 'Meghalaya', value: generateNumber(newjsondata["Meghalaya"]) },
      { id: 'MZ', state: 'Mizoram', value: generateNumber(newjsondata["Mizoram"]) },
      { id: 'NL', state: 'Nagaland', value: generateNumber(newjsondata["Nagaland"]) },
      { id: 'OR', state: 'Odisha', value: generateNumber(newjsondata["Odisha"]) },
      { id: 'PB', state: 'Punjab', value: generateNumber(newjsondata["Punjab"]) },
      { id: 'RJ', state: 'Rajasthan', value: generateNumber(newjsondata["Rajasthan"]) },
      { id: 'SK', state: 'Sikkim', value: generateNumber(newjsondata["Sikkim"]) },
      { id: 'TN', state: 'Tamil Nadu', value: generateNumber(newjsondata["Tamil Nadu"]) },
      { id: 'TG', state: 'Telangana', value: generateNumber(newjsondata["Telangana"]) },
      { id: 'TR', state: 'Tripura', value: generateNumber(newjsondata["Tripura"]) },
      { id: 'UT', state: 'Uttarakhand', value: generateNumber(newjsondata["Uttarakhand"]) },
      { id: 'UP', state: 'Uttar Pradesh', value: generateNumber(newjsondata["Uttar Pradesh"]) },
      { id: 'WB', state: 'West Bengal', value: generateNumber(newjsondata["West Bengal"]) },
      { id: 'AN', state: 'Andaman and Nicobar Islands', value: generateNumber(newjsondata["Andaman and Nicobar Islands"]) },
      { id: 'CH', state: 'Chandigarh', value: generateNumber(newjsondata["Chandigarh"]) },
      { id: 'DN', state: 'Dadra and Nagar Haveli', value: generateNumber(newjsondata["Dadra and Nagar Haveli"]) },
      { id: 'DD', state: 'Daman and Diu', value: generateNumber(newjsondata["Daman and Diu"]) },
      { id: 'DL', state: 'Delhi', value: generateNumber(newjsondata["Delhi"]) },
      { id: 'JK', state: 'Jammu and Kashmir', value: generateNumber(newjsondata["Jammu and Kashmir"]) },
      { id: 'LA', state: 'Ladakh', value: generateNumber(newjsondata["Ladakh"]) },
      { id: 'LD', state: 'Lakshadweep', value: generateNumber(newjsondata["Lakshadweep"]) },
      { id: 'PY', state: 'Puducherry', value: generateNumber(newjsondata["Puducherry"]) }
    ];
  };
  
  const [tooltipContent, setTooltipContent] = useState('');
  const [data, setData] = useState(getHeatMapData());

  const gradientData = {
    fromColor: COLOR_RANGE[0],
    toColor: COLOR_RANGE[COLOR_RANGE.length - 1],
    min: 0,
    max: data.reduce((max, item) => (item.value > max ? item.value : max), 0)
  };

  const colorScale = scaleQuantile()
    .domain(data.map(d => d.value))
    .range(COLOR_RANGE);

  const onMouseEnter = (geo, current = { value: 'NA' }) => {
    return () => {
      setTooltipContent(`${geo.properties.name}: ${current.value}`);
    };
  };

  const onMouseLeave = () => {
    setTooltipContent('');
  };

  const onChangeButtonClick = () => {
    setData(getHeatMapData());
  };

  useEffect(() => {
    FetchData()
  }, [])
  return (
    <div className="full-width-height container">
      <h6 className="no-margin center">Employee's Location</h6>
      <ReactTooltip>{tooltipContent}</ReactTooltip>
        <ComposableMap
          projectionConfig={PROJECTION_CONFIG}
          projection="geoMercator"
          width={600}
          height={220}
          data-tip=""
        >
          <Geographies geography={INDIA_TOPO_JSON}>
            {({ geographies }) =>
              geographies.map(geo => {
                //console.log(geo.id);
                const current = data.find(s => s.id === geo.id);
                return (
                  <Geography
                    key={geo.rsmKey}
                    geography={geo}
                    fill={current ? colorScale(current.value) : DEFAULT_COLOR}
                    style={geographyStyle}
                    onMouseEnter={onMouseEnter(geo, current)}
                    onMouseLeave={onMouseLeave}
                  />
                );
              })
            }
          </Geographies>
        </ComposableMap>
        <LinearGradient data={gradientData} />
        <div className="center">
          <button className="mt16" onClick={onChangeButtonClick}>Change</button>
        </div>
    </div>
  );
}

export default Map;
