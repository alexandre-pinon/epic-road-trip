describe('The Register page', () => {

    it('login with an account', () => {
        cy.visit('http://localhost:3000/login') // change URL to match your dev URL
        cy.wait(500)

        //Login
        cy.fixture('userInformations').then((userInformationsForLogin) => {
            cy.get('#email').type(userInformationsForLogin.email)
            cy.get('#password').type(userInformationsForLogin.password)

            cy.wait(1000)
            cy.get('[data-attr=login-confirm]').click()
            cy.location('pathname').should('eq', '/')
        })
    })





    it('email input only takes email', () => {
        cy.visit('http://localhost:3000/login')
        cy.wait(500)
        cy.fixture('userInformations').then((userInformationsForLogin) => {
            //fakeEmail
            cy.get('#email').type(userInformationsForLogin.fakeEmail)
            cy.wait(500)
            cy.get('#password').type("password")
        })
        cy.wait(1000)
        cy.get('#errorMessageIncorrectEmail').contains('The email is not valid')
    })




    it('check Login button is disable if all informations have not been entered', () => {
        cy.visit('http://localhost:3000/login')
        cy.wait(500)
        //login button is disable because the input are empty
        cy.get('#loginButton').should('be.disabled')
        cy.wait(500)

        //enter information to test if the login button is able now
        cy.fixture('userInformations').then((userInformationsForLogin) => {
            cy.get('#email').type(userInformationsForLogin.email)
            cy.wait(500)
            cy.get('#loginButton').should('be.disabled')
            cy.get('#password').type(userInformationsForLogin.password)
            cy.wait(500)
            cy.get('#loginButton').should('not.be.disabled') //all input are good
        })
    })





    it('check the error message if wrong credentials and clear password input', () => {
        cy.visit('http://localhost:3000/login')
        cy.wait(500)
        //login
        cy.fixture('userInformations').then((userInformationsForLogin) => {
            cy.get('#email').type(userInformationsForLogin.email)
            //wrong password
            cy.get('#passwword').type(userInformationsForLogin.wrongPassword)
            cy.wait(500)
            cy.get('#loginButton').click()
            cy.wait(1000)

            //check error message
            cy.get('#errorMessageWrongCredentials').contains('Vos identifiants ne sont pas bon, veuillez réessayer.')
            cy.wait(500)
            //check if password input are cleaned
            cy.get('#passwword').should('have.value', '');
        })
    })

})