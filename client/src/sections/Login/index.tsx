import { NumberInput, TextInput, Checkbox, Button, Group, Box, PasswordInput, Avatar, Center } from '@mantine/core';
import { useForm } from '@mantine/form';

export function Login() {
  const form = useForm({
    initialValues: {
      email: '',
      termsOfService: false,
    },

    validate: {
      email: (value) => (/^\S+@\S+$/.test(value) ? null : 'Invalid email'),
    },
  });

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
          label="Firstname"
          placeholder="Han"
          mt="sm"
        />

        <TextInput
          required
          label="Lastname"
          placeholder="SoHee"
          mt="sm"
        />

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

        <NumberInput
          required
          label="Phone"
          hideControls
          placeholder="Your phone number"
          mt="sm"
        />

        <Checkbox
          mt="sm"
          label="I agree to sell my soul"
          {...form.getInputProps('termsOfService', { type: 'checkbox' })}
        />

        <Group position="right" mt="sm">
          <Button type="submit">Submit</Button>
        </Group>
      </form>
    </Box>
  );
}
