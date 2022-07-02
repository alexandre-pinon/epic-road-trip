import { useNavigate } from 'react-router-dom';
import { TextInput, Button, Group, Box, PasswordInput, Avatar, Center } from '@mantine/core';
import { useForm } from '@mantine/form';
import axios from 'axios';
import { useState } from 'react';

export function Login() {
  const navigate = useNavigate();

  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')


  const submitDataLogin = (event: any) => {
    axios.defaults.withCredentials = true
    event.preventDefault()
    let params = {
      email: email,
      password: password,
    };
    console.log("event: ", params);

    axios({
      method: 'post',
      url: 'http://localhost:8000/api/v1/auth/login',
      data: params,
    })
      .catch(function (error) {
        console.log(error);
      });
  };

  // const form = useForm({
  //   initialValues: {
  //     email: '',
  //     termsOfService: false,
  //   },

  //   validate: {
  //     email: (value) => (/^\S+@\S+$/.test(value) ? null : 'Invalid email'),
  //   },
  // });

  // const goToHome = async () => {
  //   console.log("Dummy login!")
  //   navigate('/');
  // };

  // const goToRegister = async () => {
  //   navigate('/register');
  // };

  return (
    // <Box sx={{ maxWidth: 300 }} mx="auto">
    //   <form onSubmit={form.onSubmit((values) => console.log(values))}>

    //     <Center>
    //       <Avatar
    //         component="a"
    //         target="_blank"
    //         src="https://cf.creatrip.com/original/blog/8482/ugu92lsquro327caty8re397qbb0cdb8.png"
    //         alt="it's me"
    //         size="lg"
    //         mt="sm"
    //       />
    //     </Center>

    //     <TextInput
    //       data-attr="login-email"
    //       required
    //       label="Email"
    //       placeholder="hansohee@gmail.com"
    //       {...form.getInputProps('email')}
    //       mt="sm"
    //     />

    //     <PasswordInput
    //       data-attr="login-password"
    //       placeholder="Password"
    //       label="Password"
    //       description="Password must include at least one letter, number and special character"
    //       required
    //       mt="sm"

    //     />

    //     <Group position="right" mt="sm">
    //       <Button variant="default" radius="xl" size="sm" compact onClick={goToRegister} type="submit">Sign up</Button>
    //       <Button variant="default" radius="xl" size="sm" compact data-attr="login-confirm" onClick={goToHome} type="submit">Sign in</Button>
    //     </Group>
    //   </form>
    // </Box>

    <form onSubmit={submitDataLogin}>
      <h1 className="h3 mb-3 fw-normal">Please login</h1>

      <input type="email" className="form-control" placeholder="Email address" required
        onChange={e => setEmail(e.target.value)}
      />

      <input type="password" className="form-control" placeholder="Password" required
        onChange={e => setPassword(e.target.value)}
      />

      <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
    </form>
  );
}


