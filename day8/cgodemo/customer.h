#ifndef _CUSTOMER_H
#define _CUSTOMER_H

struct DOB{

int day;
int month;
int year;

}
struct Customer {
    const char *name;
    int customerId;
    DOB dob;

};
int displayCustomer(struct Customer *customer, char *out);

#endif