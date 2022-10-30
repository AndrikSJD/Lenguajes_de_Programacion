class Op:

    op = "";

    def __init__(self,operador) :
        self.op = operador
    
    def get_op(self):
        return self.op
    
    def __repr__(self):
        return str(self.op)