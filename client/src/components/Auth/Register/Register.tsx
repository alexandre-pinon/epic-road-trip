import { useNavigate } from 'react-router-dom';
import { NumberInput, TextInput, Checkbox, Button, Group, Box, PasswordInput, Avatar, Center } from '@mantine/core';
import { useForm } from '@mantine/form';

export function Register() {
  const navigate = useNavigate();

  const form = useForm({
    initialValues: {
      firstname: '',
      lastname: '',
      email: '',
      termsOfService: false,
    },

    validate: {
      email: (value) => (/^\S+@\S+$/.test(value) ? null : 'Invalid email'),
    },
  });

  const goToLogin = async () => {
    console.log("Dummy register!")
    navigate('/login');
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
          data-attr="firstname"
          label="Firstname"
          placeholder="Han"
          mt="sm"
          {...form.getInputProps('firstname')}
        />

        <TextInput
          required
          data-attr="lastname"
          label="Lastname"
          placeholder="SoHee"
          mt="sm"
          {...form.getInputProps('lastname')}
        />

        <TextInput
          required
          data-attr="email"
          label="Email"
          placeholder="hansohee@gmail.com"
          {...form.getInputProps('email')}
          mt="sm"
        />

        <NumberInput
          required
          data-attr="Phone"
          label="Phone"
          hideControls
          placeholder="Your phone number"
          mt="sm"
        />

        <PasswordInput
          required
          data-attr="password"
          placeholder="Password"
          label="Password"
          description="Password must include at least one letter, number and special character"
          mt="sm"
        />

        <PasswordInput
          required
          data-attr="passwordCheck"
          placeholder="Confirm Password"
          label="Confirm Password"
          description="Confirm your password"
          mt="sm"
        />

        <Checkbox
          mt="sm"
          label="I agree to sell my soul"
          {...form.getInputProps('termsOfService', { type: 'checkbox' })}
        />

        <Group position="right" mt="sm">
          <Button data-attr="register-confirm" variant="default" radius="xl" size="sm" onClick={goToLogin} type="submit">Continue</Button>
        </Group>
      </form>
    </Box>
  );
}
