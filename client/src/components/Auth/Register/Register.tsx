import { Link, useNavigate } from 'react-router-dom'
import { NumberInput, Text, TextInput, Checkbox, Button, Group, Box, PasswordInput, Avatar, Center, Paper, Container, Title, Anchor } from '@mantine/core';
import { useForm } from '@mantine/form';
import axios from 'axios';
import { toast } from 'react-toastify';
import { SyntheticEvent, useState } from 'react';

export function Register() {

  const navigate = useNavigate();

  const goLogin = async () => {
    navigate('/login');
  };

  const [firstname, setFirstName] = useState('')
  const [lastname, setLastName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [phone, setPhone] = useState('')

  const submitData = (event: any) => {
    axios.defaults.withCredentials = true
    event.preventDefault()

    let params = {
      firstName: firstname,
      lastName: lastname,
      email: email,
      password: password,
      phone: phone,
    };
    console.log(params);
    axios({
      method: 'post',
      url: 'http://localhost:8000/api/v1/auth/register',
      data: params
    });

    navigate('/login');
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

    <Container size={420} my={40}>
      <Title
        align="center"
        sx={(theme) => ({ fontFamily: `Greycliff CF, ${theme.fontFamily}`, fontWeight: 900 })}
        data-testid="welcome"
      >
        Welcome 🍩!
      </Title>

      <Text color="dimmed" size="sm" align="center" data-testid="loginPage" mt={5}>
        Do you already have an account :{' '}
        <Anchor<'a'> href="#" size="sm" onClick={goLogin}>
          Log to your account
        </Anchor>
      </Text>

      <Paper withBorder shadow="md" p={30} mt={30} radius="md">
        <form onSubmit={submitData}>

          <TextInput
            label="Firstname"
            placeholder="Your firstname"
            required
            onChange={e => setFirstName(e.target.value)}
            data-testid="Firstname"
          />
          <TextInput
            label="Lastname"
            placeholder="Your lastname"
            required
            onChange={e => setLastName(e.target.value)}
            data-testid="Lastname"
          />
          <TextInput
            label="Email"
            placeholder="Your email"
            required
            onChange={e => setEmail(e.target.value)}
            data-testid="email"
          />
          <PasswordInput
            label="Password"
            placeholder="Your password"
            required mt="md"
            onChange={e => setPassword(e.target.value)}
            data-testid="password"
          />
          <TextInput
            label="Phone"
            placeholder="Your phone"
            required
            onChange={e => setPhone(e.target.value)}
            data-testid="phone"
          />
          <Button fullWidth mt="xl" type="submit" data-testid="signUp">
            Sign up
          </Button>
        </form>
      </Paper>
    </Container >
  );
}
