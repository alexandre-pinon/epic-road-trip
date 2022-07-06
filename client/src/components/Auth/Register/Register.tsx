import { Link, useNavigate } from 'react-router-dom'
import { NumberInput, TextInput, Checkbox, Button, Group, Box, PasswordInput, Avatar, Center } from '@mantine/core';
import { useForm } from '@mantine/form';
import axios from 'axios';
import { toast } from 'react-toastify';
import { SyntheticEvent, useState } from 'react';

export function Register() {
  const navigate = useNavigate();

  const [firstName, setFirstName] = useState('')
  const [lastName, setLastName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [phone, setPhone] = useState('')

  const submitData = (data: any) => {
    let params = {
      firstname: data.firstname,
      lastname: data.lastname,
      email: data.email,
      password: data.password,
      confirmpassword: data.cpassword,
    };
    console.log(data);
    axios({
      method: 'post',
      url: 'http://localhost:8000/api/v1/auth/register',
      data: params
    });
  };




  // const form = useForm({
  //   initialValues: {
  //     firstname: '',
  //     lastname: '',
  //     email: '',
  //     password: '',
  //     phone: ''
  //   },

  //   validate: {
  //     email: (value) => (/^\S+@\S+$/.test(value) ? null : 'Invalid email'),
  //   },
  // });

  // const submitRegister = (data: any) => {
  //   axios
  //     .post("http://localhost:8000/api/v1/auth/register")
  //     .then(function (response) {
  //       toast.success(response.data.message, {
  //         position: "top-right",
  //         autoClose: 3000,
  //         hideProgressBar: true,
  //         closeOnClick: true,
  //         pauseOnHover: true,
  //         draggable: false,
  //         progress: 0,
  //         toastId: "my_toast",
  //       });
  //     })
  //     .catch(function (error) {
  //       console.log(error);
  //     });
  // };

  // const goToLogin = async () => {
  //   console.log("Dummy register!")
  //   navigate('/login');
  // };

  return (
    <form onSubmit={submitData}>
      <h1 className="h3 mb-3 fw-normal">Please register</h1>

      <input className="form-control" placeholder="Name" required
        onChange={e => setFirstName(e.target.value)}
      />

      <input className="form-control" placeholder="Name" required
        onChange={e => setLastName(e.target.value)}
      />

      <input type="email" className="form-control" placeholder="Email address" required
        onChange={e => setEmail(e.target.value)}
      />

      <input type="password" className="form-control" placeholder="Password" required
        onChange={e => setPassword(e.target.value)}
      />

      <input className="form-control" placeholder="Phone" required
        onChange={e => setPhone(e.target.value)}
      />

      <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
    </form>
  );
}
