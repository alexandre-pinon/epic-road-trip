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
        cy.wait(2000)
        cy.location('href').should('include', '/login#')
    })
})