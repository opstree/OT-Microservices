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
  '#9a311f',
  '#CCCCFF'
];

const DEFAULT_COLOR = '	#92A1CF';

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

// will generate random heatmap data on every call

function Map() {
  const [stats, handleStats] = useState([]);

  const FetchData = async () => {
    const jsondata = await fetch('http://172.17.0.3:8080/search/city');
    const stats = await jsondata.json();
    handleStats(stats)  
  }
  const newjsondata = stats
  console.log(newjsondata["Uttarakhand"])

  const getHeatMapData = () => {
    return [
      { id: 'AP', state: 'Andhra Pradesh', value: newjsondata["Andhra Pradesh"] },
      { id: 'AR', state: 'Arunachal Pradesh', value: newjsondata["Arunachal Pradesh"] },
      { id: 'AS', state: 'Assam', value: newjsondata["Assam"] },
      { id: 'BR', state: 'Bihar', value: newjsondata["Bihar"] },
      { id: 'CT', state: 'Chhattisgarh', value: newjsondata["Chhattisgarh"] },
      { id: 'GA', state: 'Goa', value: newjsondata["Goa"] },
      { id: 'GJ', state: 'Gujarat', value: newjsondata["Gujarat"] },
      { id: 'HR', state: 'Haryana', value: newjsondata["Haryana"] },
      { id: 'HP', state: 'Himachal Pradesh', value: newjsondata["Himachal Pradesh"] },
      { id: 'JH', state: 'Jharkhand', value: newjsondata["Jharkhand"] },
      { id: 'KA', state: 'Karnataka', value: newjsondata["Karnataka"] },
      { id: 'KL', state: 'Kerala', value: newjsondata["Kerala"] },
      { id: 'MP', state: 'Madhya Pradesh', value: newjsondata["Madhya Pradesh"] },
      { id: 'MH', state: 'Maharashtra', value: newjsondata["Maharashtra"] },
      { id: 'MN', state: 'Manipur', value: newjsondata["Manipur"] },
      { id: 'ML', state: 'Meghalaya', value: newjsondata["Meghalaya"] },
      { id: 'MZ', state: 'Mizoram', value: newjsondata["Mizoram"] },
      { id: 'NL', state: 'Nagaland', value: newjsondata["Nagaland"] },
      { id: 'OR', state: 'Odisha', value: newjsondata["Odisha"] },
      { id: 'PB', state: 'Punjab', value: newjsondata["Punjab"] },
      { id: 'RJ', state: 'Rajasthan', value: newjsondata["Rajasthan"] },
      { id: 'SK', state: 'Sikkim', value: newjsondata["Sikkim"] },
      { id: 'TN', state: 'Tamil Nadu', value: newjsondata["Tamil Nadu"] },
      { id: 'TG', state: 'Telangana', value: newjsondata["Telangana"] },
      { id: 'TR', state: 'Tripura', value: newjsondata["Tripura"] },
      { id: 'UT', state: 'Uttarakhand', value: newjsondata["Uttarakhand"] },
      { id: 'UP', state: 'Uttar Pradesh', value: newjsondata["Uttar Pradesh"] },
      { id: 'WB', state: 'West Bengal', value: newjsondata["West Bengal"] },
      { id: 'AN', state: 'Andaman and Nicobar Islands', value: newjsondata["Andaman and Nicobar Islands"] },
      { id: 'CH', state: 'Chandigarh', value: newjsondata["Chandigarh"] },
      { id: 'DN', state: 'Dadra and Nagar Haveli', value: newjsondata["Dadra and Nagar Haveli"] },
      { id: 'DD', state: 'Daman and Diu', value: newjsondata["Daman and Diu"] },
      { id: 'DL', state: 'Delhi', value: newjsondata["Delhi"] },
      { id: 'JK', state: 'Jammu and Kashmir', value: newjsondata["Jammu and Kashmir"] },
      { id: 'LA', state: 'Ladakh', value: newjsondata["Ladakh"] },
      { id: 'LD', state: 'Lakshadweep', value: newjsondata["Lakshadweep"] },
      { id: 'PY', state: 'Puducherry', value: newjsondata["Puducherry"] }
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
