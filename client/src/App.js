import {Col,Container,Row} from "react-bootstrap";

import axios from "axios";
import { useState , useEffect } from "react";

import "./App.css"

function App() {

  const [apiData,setApiData] = useState(false);

  useEffect(() => {
   
    const fetchdata = async() => {
     // const apiUrl = process.env.REACT_APP_API_ROOT;
      const apiUrl  = "http://localhost:8000";
     
      try {

      const response = await axios.get(apiUrl);

      if(response.status === 200){
        if(response?.data.status==="ok"){
          setApiData(response?.data?.blog_records);
        }
        
      }
      
    } catch(error){
        console.log(error.response);
      }

       
    }
  
    fetchdata();

    return () => {}
  }, []);
  
  console.log(apiData);
  return (
  <Container>
    <Row>
      <Col xs="12" className="py-2">
        <h1 className="text-center">
            React application with fiber backend
        
        </h1>
      </Col>
     {/* {
  apiData &&
  apiData.map((record, index) => (
    <Col xs="4" className="py-5 box" key={index}>
      console.log(record);
      <div className="title">{record.title}</div>
     
    </Col>
  ))
} */}

{apiData.length > 0 ? (
  apiData.map((record, index) => (
    <Col xs="{4}" className="py-5 box" key={index}>
      <div className="title">{record.title}</div>
    </Col>
  ))
) : (
  <p>No data found</p>
)}
    </Row>
  </Container>
  );
}

export default App;
