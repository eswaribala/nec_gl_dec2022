#include "customer.h"
#include <stdio.h>

int displayCustomer(struct Customer *customer, char *out) {
    int n;

    n = sprintf(out, "Customer details, %s from %d/%d/%d! We come in peace :)", customer->name, customer->dob.day,customer->dob.month,customer->dob.year);

    return n;
}