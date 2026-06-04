import axios from 'axios';
import React, { useState ,useEffect} from 'react'
import { useParams } from 'react-router-dom';
import { Col, Container, Row } from "react-bootstrap";


const Blog = () =>{

  const params = useParams();
  const [apiData, setApiData] = useState([]);

  useEffect(() => {
      const fetchData = async () => {
        const apiUrl = "http://localhost:8000/"+params.id;
  
        try {
          const response = await axios.get(apiUrl);
  
          if (
            response.status === 200 &&
            response.data.status === "ok"
          ) {
            setApiData(response.data.record);
          }
        } catch (error) {
          console.log(error);
        }
      };
  
      fetchData();
    }, []);

    console.table(apiData);
    return (
    <Container>
      <Row>
        <Col xs="12"> <h1>{apiData.title}</h1></Col>
        <Col xs="12"> {apiData.post}</Col>
      </Row>
    </Container>
  )
}

export default Blog