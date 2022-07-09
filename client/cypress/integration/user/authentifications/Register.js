describe('Register page', () => {

    it('Check attributes Register page', () => {
        cy.visit('http://localhost:3000/register')

        //Firstname
        cy.get('[data-testid="Firstname"]')
            .should('have.attr', 'required')

        cy.get('[data-testid="Firstname"]')
            .should('have.attr', 'placeholder')
            .and('equal', 'Your firstname')

        //Lastname
        cy.get('[data-testid="Lastname"]')
            .should('have.attr', 'required')

        cy.get('[data-testid="Lastname"]')
            .should('have.attr', 'placeholder')
            .and('equal', 'Your lastname')

        //email
        cy.get('[data-testid="email"]')
            .should('have.attr', 'required')

        cy.get('[data-testid="email"]')
            .should('have.attr', 'placeholder')
            .and('equal', 'Your email')

        //password
        cy.get('[data-testid="password"]')
            .should('have.attr', 'required')

        cy.get('[data-testid="password"]')
            .should('have.attr', 'placeholder')
            .and('equal', 'Your password')

        //phone
        cy.get('[data-testid="phone"]')
            .should('have.attr', 'required')

        cy.get('[data-testid="phone"]')
            .should('have.attr', 'placeholder')
            .and('equal', 'Your phone')
    })



    it('check text on the Register page', () => {
        cy.visit('http://localhost:3000/register')

        cy.get('[data-testid="welcome"]').contains('Welcome 🍩!')

        cy.get('[data-testid="loginPage"]')
            .contains('Do you already have an account : Log to your account')


        cy.get('form > :nth-child(1) > :nth-child(1)').contains('Firstname')
        cy.get('form > :nth-child(2) > :nth-child(1)').contains('Lastname')
        cy.get('form > :nth-child(3) > :nth-child(1)').contains('Email')
        cy.get('form > :nth-child(4) > :nth-child(1)').contains('Password')
        cy.get('form > :nth-child(5) > :nth-child(1)').contains('Phone')

        cy.get('[data-testid="signUp"]').contains('Sign up')
    })



    it('check Login button Page from Register page', () => {
        cy.visit('http://localhost:3000/register')

        cy.get('[data-testid="loginPage"] > .mantine-Text-root').click()
        cy.wait(500)
        cy.location('href').should('include', '/login#')
    })



    it('check creation account Register page', () => {
        cy.visit('http://localhost:3000/register')

        cy.get('[data-testid="Firstname"]').type('Oran')
        cy.get('[data-testid="Firstname"]').should('have.value', 'Oran')
        cy.get('[data-testid="Lastname"]').type('Gina')
        cy.get('[data-testid="Lastname"]').should('have.value', 'Gina')
        cy.get('[data-testid="email"]').type('orangina@mail.fr')
        cy.get('[data-testid="email"]').should('have.value', 'orangina@mail.fr')
        cy.get('[data-testid="password"]').type('password')
        cy.get('[data-testid="password"]').should('have.value', 'password')
        cy.get('[data-testid="phone"]').type('0632437596')
        cy.get('[data-testid="phone"]').should('have.value', '0632437596')

    })


})