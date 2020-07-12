
struct ListNode 
{
    int value;
    ListNode * next;
};


ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
 
 
 ListNode* h = new(ListNode);
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

 }while(l1 != nullptr && l2 != nullptr)

if (l1 != nullptr) {
    cur->next = l1;
}

if (l2 != nullptr) {
    cur->next = l2;
}

return h->next;
}