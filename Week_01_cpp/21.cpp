ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
 
 
 ListNode* h = New(ListNode);
 ListNode* cur = h;
 

 do {

    if (l1->value <= l2->value) {
        //插入q
        cur->next = l1; 
        l1 = l1->next;
    } else {
        cur->next = l2;
        l2 = l2->next;
    }
    cur = cur->next;

 }while(l1 != null && l2 != null)

if (l1->next != null) {
    cur->next = l1;
}

if (l2->next != null) {
    cur->next = l2;
}

return h->next;
}