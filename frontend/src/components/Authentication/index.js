import { Form, Table, Button, Modal } from "react-bootstrap";
import { useState, createRef } from "react";
import user from "../../data/user.json"

export default function Authentication() {
  const [users, setUser] = useState(user);
  const [status, setStatus] = useState("");
  const [message, setMessage] = useState("");
  const [showModal, setShowModal] = useState(false);
  const [validated, setValidated] = useState(false);

  const formData = createRef();

  const handleReset = () => {
    formData.current.reset();
    setValidated(false);
  };

  const add = (event) => {
    event.preventDefault();

    if (users.map((e) => e.email).includes(formData.current.email.value)) {
      setStatus("Error");
      setMessage("Email already exist");
      setShowModal(true);
      return;
    }

    if (!formData.current.email.value || !formData.current.password.value) {
      setStatus("Error");
      setMessage("Please fill out all required fields");
      setShowModal(true);
      setValidated(false);
      return;
    }

    const newUser = {
      email: formData.current.email.value,
      password: formData.current.password.value,
    };

    setUser([...users, newUser]);
    setValidated(true);
    setStatus("Success");
    setMessage("Successfully add user");
    setShowModal(true);
    handleReset();
  };

  const handleClose = () => setShowModal(false);

  return (
    <div>
      <Modal show={showModal} onHide={handleClose}>
        <Modal.Header closeButton>
          <Modal.Title>{status}</Modal.Title>
        </Modal.Header>
        <Modal.Body>{message}</Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Close
          </Button>
        </Modal.Footer>
      </Modal>

      <div
        style={{
          alignItems: "center",
          justifyContent: "center",
          padding: "20px",
        }}
      >
        <Form onSubmit={add} validated={validated} ref={formData}>
          <Form.Group>
            <Form.Label>Email:</Form.Label>
            <Form.Control
              required="true"
              type="email"
              placeholder="johndoe@mail.com"
              name="email"
            />

            <Form.Label>Password:</Form.Label>
            <Form.Control
              required="true"
              type="text"
              placeholder="abcd2134!"
              name="password"
            />
          </Form.Group>
          <br />
          <Button variant="primary" type="submit">
            Add New Auth
          </Button>
        </Form>
      </div>
      <Table striped bordered hover responsive variant="dark">
        <thead>
          <tr>
            <th>No</th>
            <th>Email</th>
            <th>Password</th>
          </tr>
        </thead>
        <tbody>
          {users.map((item, index) => {
            return (
              <tr key={index}>
                <td>{index + 1}</td>
                <td>{item.email}</td>
                <td>{item.password}</td>
              </tr>
            );
          })}
        </tbody>
      </Table>
    </div>
  );
}
