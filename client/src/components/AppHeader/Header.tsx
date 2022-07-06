import { useNavigate } from 'react-router-dom';
import { createStyles, AppShell, Button, Container, Group, Header, Title, ActionIcon, Image, Grid, Text } from '@mantine/core';
import { BrandGithub, BrandInstagram, BrandYoutube } from 'tabler-icons-react';
import logo from "./assets/one_piece_wiki_logo.png";
import axios from 'axios';

const logout = (event: any) => {
  axios.defaults.withCredentials = true
  event.preventDefault()

  axios({
    method: 'post',
    url: 'http://localhost:8000/api/v1/auth/logout',
  })
    .catch(function (error) {
      console.log(error);
    });
};

const useStyles = createStyles((theme) => ({
  button: {
    // backgroundColor: '#424242',
    // border: 0,
    height: 42,
    paddingLeft: 20,
    paddingRight: 20,

    '&:hover': {
      backgroundColor: theme.fn.darken('#B3E5FC', 0.05),
    },
  },

  leftIcon: {
    marginRight: 15,
  },

  social: {
    width: 260,

    [theme.fn.smallerThan('sm')]: {
      width: 'auto',
      marginLeft: 'auto',
    },
  },
}));

export function AppHeader() {
  const { classes } = useStyles();
  const navigate = useNavigate();

  const goToLogin = async () => {
    console.log("Dummy register!")
    navigate('/login');
  };

  const goToHome = async () => {
    navigate('/');
  };

  return (
    <AppShell
      padding="md"

      header={
        <Container>
          <Header height={100} p="xs">
            <Group sx={{ height: '100%' }} px={20} position='apart'>

              <Title
                order={4}
              >
                <Grid>
                  <Grid.Col span={12}>
                    <Image height={50} src={logo} onClick={goToHome} />
                  </Grid.Col>

                </Grid>
                <Text
                  align="center"
                  variant="gradient"
                  gradient={{ from: '#616161', to: '#B3E5FC', deg: 90 }}
                >
                  Epic Road Trip
                </Text>

              </Title>
              <Group spacing={0} className={classes.social} position="center" noWrap>
                <ActionIcon size="lg">
                  <a target="_blank" rel="noreferrer" href="https://github.com/alexandre-pinon/epic-road-trip">
                    <BrandGithub color="#616161" size={20} />
                  </a>
                </ActionIcon>
                <ActionIcon size="lg">
                  <a target="_blank" rel="noreferrer" href="https://www.youtube.com/watch?v=GHBLNXXdZ3c">
                    <BrandYoutube color="#616161" size={20} />
                  </a>
                </ActionIcon>
                <ActionIcon size="lg">
                  <a target="_blank" rel="noreferrer" href="https://www.instagram.com/onepiece_staff/?hl=fr">
                    <BrandInstagram color="#616161" size={20} />
                  </a>
                </ActionIcon>
              </Group>
              <Group>
                <Button
                  radius={50}
                  className={classes.button}
                  onClick={goToLogin}
                  variant="default"
                >
                  Sign in
                </Button>
                <Button
                  radius={50}
                  className={classes.button}
                  onClick={logout}
                  variant="default"
                >
                  Logout
                </Button>
              </Group>
            </Group>
          </Header>
        </ Container>}
    >
    </AppShell>
  );
}