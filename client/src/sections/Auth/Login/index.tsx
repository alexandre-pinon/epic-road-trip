import { useNavigate } from 'react-router-dom';
import { TextInput, Button, Group, Box, PasswordInput, Avatar, Center } from '@mantine/core';
import { useForm } from '@mantine/form';

export function Login() {
  const navigate = useNavigate();

  const form = useForm({
    initialValues: {
      email: '',
      termsOfService: false,
    },

    validate: {
      email: (value) => (/^\S+@\S+$/.test(value) ? null : 'Invalid email'),
    },
  });

  const goToHome = async () => {
    console.log("Dummy login!")
    navigate('/');
  };

  const goToRegister = async () => {
    navigate('/register');
  };

  return (
    <Box sx={{ maxWidth: 300 }} mx="auto">
      <form onSubmit={form.onSubmit((values) => console.log(values))}>

        <Center>
          <Avatar
            component="a"
            target="_blank"
            src="https://cf.creatrip.com/original/blog/8482/ugu92lsquro327caty8re397qbb0cdb8.png"
            alt="it's me"
            size="lg"
            mt="sm"
          />
        </Center>

        <TextInput
          required
          label="Email"
          placeholder="hansohee@gmail.com"
          {...form.getInputProps('email')}
          mt="sm"
        />

        <PasswordInput
          placeholder="Password"
          label="Password"
          description="Password must include at least one letter, number and special character"
          required
          mt="sm"

        />

        <Group position="right" mt="sm">
          <Button onClick={goToRegister} type="submit">Sign up</Button>
          <Button onClick={goToHome} type="submit">Sign in</Button>
        </Group>
      </form>
    </Box>
  );
}
