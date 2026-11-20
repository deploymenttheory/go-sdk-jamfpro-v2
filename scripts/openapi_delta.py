import json,sys
old,new=map(lambda p:json.load(open(p)),sys.argv[1:3])
def clean(x):
 if isinstance(x,dict):return {k:clean(v) for k,v in x.items() if k not in ['description','summary','example','examples','externalDocs','operationId','tags']}
 if isinstance(x,list):return [clean(v) for v in x]
 return x
for section,a,b in [('paths',old['paths'],new['paths']),('schemas',old['components']['schemas'],new['components']['schemas'])]:
 print('\n'+section)
 for k in sorted(a.keys()|b.keys()):
  if k not in a: print('ADDED',k)
  elif k not in b:print('REMOVED',k)
  elif clean(a[k])!=clean(b[k]):
   print('CHANGED',k)
   import difflib
   print('\n'.join(difflib.unified_diff(json.dumps(clean(a[k]),indent=2).splitlines(),json.dumps(clean(b[k]),indent=2).splitlines(),n=2)))
