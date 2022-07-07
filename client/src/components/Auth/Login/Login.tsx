import { useNavigate } from 'react-router-dom';
import {
  Text,
  TextInput,
  Button,
  Group,
  Box,
  PasswordInput,
  Avatar,
  Center,
  Paper,
  PaperProps,
  Divider,
  Checkbox,
  Anchor,
  Container,
  Title,
} from '@mantine/core';
// import { useForm } from '@mantine/form';
import { useForm, useToggle, upperFirst } from '@mantine/hooks';

import axios from 'axios';
import { useState } from 'react';


export function Login({ auth }: any) {
  const [type, toggle] = useToggle('login', ['login', 'register']);
  const navigate = useNavigate();

  const goRegister = async () => {
    navigate('/register');
  };

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
      .then((response) => {
        console.log(response.data);
        console.log("[LOGIN] auth.setAuthenticated current: ", auth.setAuthenticated);
        auth.setIsAuthenticated(true);
        console.log("[LOGIN] auth.setAuthenticated after: ", auth.setAuthenticated);
        navigate('/');
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

    <Container size={420} my={40}>
      <Title
        align="center"
        sx={(theme) => ({ fontFamily: `Greycliff CF, ${theme.fontFamily}`, fontWeight: 900 })}
      >
        Welcome 👋!
      </Title>

      <Text color="dimmed" size="sm" align="center" mt={5}>
        Do not have an account yet?{' '}
        <Anchor<'a'> href="#" size="sm" onClick={goRegister}>
          Create account
        </Anchor>
      </Text>

      <Paper withBorder shadow="md" p={30} mt={30} radius="md">
        <form onSubmit={submitDataLogin}>

          <TextInput
            label="Email"
            placeholder="Your email"
            required
            onChange={e => setEmail(e.target.value)}
          />
          <PasswordInput
            label="Password"
            placeholder="Your password"
            required mt="md"
            onChange={e => setPassword(e.target.value)}
          />
          <Button fullWidth mt="xl" type="submit">
            Sign in
          </Button>
        </form>

      </Paper>
    </Container>


  );
}


