import axios from "axios";
import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { Button, Container, Form, Row, Col } from "react-bootstrap";


function EditBlog(){

    const params = useParams();
    const navigate = useNavigate();

      const [formData, setFormData] = useState({
        title: "",
        post: "",
    });

  /*
    useEffect(()=>{
        getBlog();
    },[params.id]);

    const getBlog = async() =>{
        try{
            
            const apiUrl = "http://localhost:8000/blog/show/"+params.id;

            const response = await axios.get(apiUrl);

            setFormData({
            title: response.data.record.title,
            post: response.data.record.post,
        });

        }catch(error){
            console.log(error)
        }
    }
*/

useEffect(() => {
    const getBlog = async () => {
        try {
            const apiUrl = `http://localhost:8000/blog/show/${params.id}`;
            const response = await axios.get(apiUrl);

            setFormData({
                title: response.data.record.title,
                post: response.data.record.post,
            });
        } catch (error) {
            console.log(error);
        }
    };

    getBlog();
}, [params.id]);
    
    const [errors,setErrors] = useState({});

    const handleChange = (e) => {
        setFormData({
            ...formData,
            [e.target.name] : e.target.value,
        })
    }

    const submitEditBlog = async(e) => {

        e.preventDefault();

        try {
            const response = await axios.post(
                "http://localhost:8000/blog/edit/"+params.id,
                formData
            );

            if (response.data.status === "ok"){
                navigate("/");
            }
        } catch (error) {
            if(error.response?.data?.errors){
                setErrors(error.response.data.errors);
            }
        }
    };


    return(
        <Container>
            <Row ClassName="justify-content-center">
                <Col md={6}>
                    <h2 ClassName="my-4">Edit Blog</h2>

                    <Form onSubmit={submitEditBlog}>

                     <Form.Control
    type="text"
    name="title"
    value={formData.title}
    onChange={handleChange}
/>

{errors.title && (
    <small className="text-danger">
        {errors.title}
    </small>
)}


<Form.Control
    as="textarea"
    rows={5}
    name="post"
    value={formData.post}
    onChange={handleChange}
/>

{errors.post && (
    <small className="text-danger">
        {errors.post}
    </small>
)}

                        <Button type="submit">
                            Save Edit Blog
                        </Button>
                    </Form>
                </Col>
            </Row>
        </Container>
    );
}


export default EditBlog;